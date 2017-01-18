package pg

import (
	"encoding/binary"
	"fmt"
	"hash/fnv"
)

// DbKey represents a postgres lock key
type DbKey int64

// NewDbKey takes an arbitrary set of ints and returns
// a single dbkey suitable for postgres advisory locking.
// This allows for additional application level namespacing, beyon
// what postgres offers.
// As with any hashing, there are the potentital for collisions
// so understand what a false positive lock behavior will cause
// in the particular scenario.
func NewDbKey(keys ...int64) DbKey {
	h := fnv.New64a()
	err := binary.Write(h, binary.LittleEndian, keys)
	if err != nil {
		panic(fmt.Sprintf("DBKey invalid keys: %s", err))
	}

	return DbKey(int64(h.Sum64()))
}
