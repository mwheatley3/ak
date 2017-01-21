package db

import (
	"github.com/mwheatley3/ak/server/pg"
)

// CreateUsersTable creates a new users table if it does not exist
func (db *Db) CreateUsersTable() error {
	var (
		r Count
		p = pg.NewParams("users")
		q = `SELECT count(*) FROM information_schema.tables WHERE table_schema = 'public' AND table_name = $1`
	)

	err := db.db.Get(&r, q, p)

	if r.Count == 0 {
		db.db.Exec(`DROP TABLE users`, nil)
		db.db.Exec(`DROP DOMAIN password_type`, nil)
		db.db.Exec(`CREATE DOMAIN password_type as TEXT CHECK (VALUE IN ('bcrypt'))`, nil)
		db.db.Exec(`
    				CREATE TABLE users (
    					id UUID NOT NULL,
    					email TEXT NOT NULL,
    					hashed_password BYTEA NOT NULL,
    					password_type password_type NOT NULL DEFAULT 'bcrypt',
    					auth_token TEXT NOT NULL,
    					system_admin BOOL NOT NULL DEFAULT FALSE,
    					created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    					updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    					deleted_at TIMESTAMPTZ DEFAULT NULL,

    					PRIMARY KEY (id)
    				)
    			`, nil)
		db.db.Exec(`CREATE UNIQUE INDEX ON users (email) WHERE deleted_at IS NULL`, nil)
		db.db.Exec(`CREATE UNIQUE INDEX ON users (auth_token) WHERE deleted_at IS NULL`, nil)
	}

	if err != nil {
		return db.unknownErr(err)
	}

	return nil
}
