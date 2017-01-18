package pg

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"sync"
	"time"
)

type tx struct {
	*Db
	db         *sqlx.DB
	mu         sync.Mutex
	depth      int
	rolledBack bool
}

func (tx *tx) Begin() error {
	tx.mu.Lock()
	defer tx.mu.Unlock()

	if tx.depth > 0 {
		tx.depth++
		return nil
	}

	txn, err := tx.db.Beginx()

	if err != nil {
		return err
	}

	tx.depth++

	tx.Db.queryer = queryer{
		Logger:        tx.Logger,
		Queryer:       txn,
		Execer:        txn,
		Preparer:      txn,
		slowThreshold: time.Duration(tx.conf.SlowThreshold) * time.Millisecond,
	}

	return nil
}

func (tx *tx) Commit() error {
	tx.mu.Lock()
	defer func() {
		tx.depth--
		tx.mu.Unlock()
	}()

	// if we're in a nested transaction
	// just move on
	if tx.depth > 1 {
		return nil
	}

	return tx.queryer.Queryer.(*sqlx.Tx).Commit()
}

func (tx *tx) Rollback() error {
	tx.mu.Lock()
	defer func() {
		tx.depth--
		tx.mu.Unlock()
	}()

	if tx.rolledBack {
		return nil
	}

	if err := tx.queryer.Queryer.(*sqlx.Tx).Rollback(); err != nil {
		return err
	}

	tx.rolledBack = true

	return nil
}

// Run runs the provided function inside a transaction
func (tx *tx) Run(fn func(*Db) error) (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			// Don't try to rollback twice (if we're panicing after trying to rollback already)
			// attempt to rollback transaction
			err = tx.Rollback()
			if err != nil {
				// propagate panic
				panic(fmt.Sprintf("Rollback error: "+err.Error()+"\nOriginal Panic: %s", panicErr))
			}

			// propagate original error
			panic(panicErr)
		}

		if err != nil {
			err2 := tx.Rollback()

			if err2 != nil {
				panic("Rollback error: " + err2.Error() + "\nOriginal Error: " + err.Error())
			}
			return
		}

		err = tx.Commit()
		return
	}()

	err = fn(tx.Db)

	return
}

// WithLockTxn attempts to request mutually exclusive lock transaction and call lockSuccess if it's obtained
// or will call lockFailure immediately if it can't be obtained
func (tx *tx) WithLockTxn(key DbKey, lockSuccess func(tx *Db) error, lockFailure func(tx *Db) error) error {
	return tx.Run(func(tx *Db) error {
		var (
			lockObtained bool
			sql          = `SELECT pg_try_advisory_xact_lock($1)`
		)

		if err := tx.Get(&lockObtained, sql, NewParams(key)); err != nil {
			return err
		}

		if lockObtained {
			return lockSuccess(tx)
		}

		if lockFailure == nil {
			return nil
		}

		return lockFailure(tx)
	})
}

// WithBlockingLockTxn requests a mutually exclusive lock transaction and blocks until it can obtain the lock
// once the lock is obtained it will call lockSuccess
func (tx *tx) WithBlockingLockTxn(key DbKey, lockSuccess func(tx *Db) error) error {
	return tx.Run(func(tx *Db) error {
		var (
			sql = `SELECT pg_advisory_xact_lock($1)`
		)

		if _, err := tx.Exec(sql, NewParams(key)); err != nil {
			return err
		}

		return lockSuccess(tx)
	})
}
