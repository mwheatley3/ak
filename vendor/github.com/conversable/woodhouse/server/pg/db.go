package pg

import (
	"crypto/tls"
	"database/sql"
	"errors"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

// Config is the config for the database connect
type Config struct {
	Host           string
	Port           uint16
	User           string
	Password       string
	Database       string
	SslMode        string
	SlowThreshold  int
	MaxConnections int
}

// NewDb returns a new DB object
func NewDb(l *logrus.Logger, conf Config) *Db {
	return &Db{Logger: l, conf: conf, once: &sync.Once{}}
}

// Db wraps a pgx.ConnPool and sqlx and provides helpful query
// and transaction functionality
type Db struct {
	*logrus.Logger
	conf Config
	queryer
	Pool         *pgx.ConnPool
	db           *sqlx.DB
	once         *sync.Once
	afterConnect []func(*pgx.Conn) error
	tx           *tx
}

// AfterConnect registers a function to run after every
// new connection is established with the db.  Use this
// for any sort of per connection init necessary (prepared statements,
// for instance)
func (d *Db) AfterConnect(fn func(*pgx.Conn) error) {
	d.afterConnect = append(d.afterConnect, fn)
}

// Close shuts down the pool
func (d *Db) Close() error {
	if d.db == nil {
		return nil
	}

	if err := d.db.Close(); err != nil {
		return err
	}

	d.Pool.Close()
	return nil
}

// Connect attempts to connect to the database
// This method is safe to run concurrently and/or multiple
// times, however it's body will only run once
func (d *Db) Connect() (err error) {
	d.once.Do(func() {
		var (
			pool   *pgx.ConnPool
			db     *sql.DB
			sqlxDB *sqlx.DB
		)

		conf := pgx.ConnConfig{
			Host:     d.conf.Host,
			Port:     d.conf.Port,
			Database: d.conf.Database,
			User:     d.conf.User,
			Password: d.conf.Password,
		}

		if err = configSSL(d.conf.SslMode, &conf); err != nil {
			return
		}

		pconf := pgx.ConnPoolConfig{
			ConnConfig:     conf,
			MaxConnections: d.conf.MaxConnections,
			AfterConnect: func(c *pgx.Conn) error {
				for _, fn := range d.afterConnect {
					if err := fn(c); err != nil {
						return err
					}
				}

				return nil
			},
		}

		pool, err = pgx.NewConnPool(pconf)

		if err != nil {
			return
		}

		db, err = stdlib.OpenFromConnPool(pool)

		if err != nil {
			return
		}

		sqlxDB = sqlx.NewDb(db, "pgx")

		d.queryer = queryer{
			Logger:        d.Logger,
			Queryer:       sqlxDB,
			Execer:        sqlxDB,
			Preparer:      sqlxDB,
			slowThreshold: time.Duration(d.conf.SlowThreshold) * time.Millisecond,
		}

		d.Pool = pool
		d.db = sqlxDB
	})

	return err
}

func (d *Db) txn() *tx {
	if d.tx != nil {
		return d.tx
	}

	tx := &tx{
		Db: &Db{
			conf:   d.conf,
			Logger: d.Logger,
			Pool:   d.Pool,
			once:   d.once,
		},
		db: d.queryer.Queryer.(*sqlx.DB),
	}

	tx.Db.tx = tx

	return tx
}

func (d *Db) startTxn() (*tx, error) {
	tx := d.txn()

	if err := tx.Begin(); err != nil {
		return nil, err
	}

	return tx, nil
}

// WithTxn runs the provided function inside a transaction
func (d *Db) WithTxn(fn func(*Db) error) (err error) {
	txn, err := d.startTxn()

	if err != nil {
		return err
	}

	return txn.Run(fn)
}

// WithLockTxn attempts to request mutually exclusive lock transaction and call lockSuccess if it's obtained
// or will call lockFailure immediately if it can't be obtained
func (d *Db) WithLockTxn(key DbKey, lockSuccess func(tx *Db) error, lockFailure func(tx *Db) error) error {
	txn, err := d.startTxn()

	if err != nil {
		return err
	}

	return txn.WithLockTxn(key, lockSuccess, lockFailure)
}

// WithBlockingLockTxn requests a mutually exclusive lock transaction and blocks until it can obtain the lock
// once the lock is obtained it will call lockSuccess
func (d *Db) WithBlockingLockTxn(key DbKey, lockSuccess func(tx *Db) error) error {
	txn, err := d.startTxn()

	if err != nil {
		return err
	}

	return txn.WithBlockingLockTxn(key, lockSuccess)
}

// taken from jackc/pgx/conn.go:511
func configSSL(sslmode string, cc *pgx.ConnConfig) error {
	// Match libpq default behavior
	if sslmode == "" {
		sslmode = "prefer"
	}

	switch sslmode {
	case "disable":
	case "allow":
		cc.UseFallbackTLS = true
		cc.FallbackTLSConfig = &tls.Config{InsecureSkipVerify: true}
	case "prefer":
		cc.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		cc.UseFallbackTLS = true
		cc.FallbackTLSConfig = nil
	case "require", "verify-ca", "verify-full":
		cc.TLSConfig = &tls.Config{
			ServerName: cc.Host,
		}
	default:
		return errors.New("sslmode is invalid")
	}

	return nil
}
