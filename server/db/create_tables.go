package db

// import (
// 	"github.com/satori/go.uuid"
// )
//
// // CreateClient creates a new Client
// func (db *Db) CreateClient(name string) (*api.Client, error) {
// 	var (
// 		v dbClient
// 		p = pg.NewParams(uuid.NewV4(), name)
// 		q = `INSERT INTO clients (id, name) VALUES ($1, $2) RETURNING ` + clientCols
// 	)
//
// 	err := db.db.Get(&v, q, p)
//
// 	if err != nil {
// 		return nil, db.unknownErr(err)
// 	}
//
// 	return v.Client, nil
// }
