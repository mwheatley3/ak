package db

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/mwheatley3/ak/server/pg"
)

// New returns a new Admin Db
func New(l *logrus.Logger, db *pg.Db) *Db {
	return &Db{
		logger: l,
		db:     db,
	}
}

// NewFromConfig initializes a pg.Db from a pg.Config before
// construction
func NewFromConfig(l *logrus.Logger, conf pg.Config) *Db {
	db := pg.NewDb(l, conf)
	return New(l, db)
}

// A Db provides database functionality for Admin functionality
type Db struct {
	logger *logrus.Logger
	db     *pg.Db
}

// Init will initialize the Db.  This is safe to be called multiple
// times and should be called before any other Service methods are used
func (db *Db) Init() error {
	return db.db.Connect()
}

func (db *Db) unknownErr(err error) error {
	db.logger.Errorf("Unhandled error, returning as ErrUnknownError: %#v", err)
	return errors.New("unknown DB error")
}
