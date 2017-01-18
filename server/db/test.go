package db

import (
	"github.com/mwheatley3/ak/server/pg"
	// "github.com/satori/go.uuid"
)

// Test creates a new Client
func (db *Db) Test(name string) (*User, error) {
	var (
		v User
		// p = pg.NewParams(uuid.NewV4(), name)
		p = pg.NewParams("hi")
		q = `SELECT * FROM test WHERE email = $1`
	)

	err := db.db.Get(&v, q, p)

	if err != nil {
		return nil, db.unknownErr(err)
	}

	return &v, nil
}
