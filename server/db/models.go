package db

import (
	// "encoding/json"

	"github.com/satori/go.uuid"
	// "time"
)

// A User is a person who can log into the admin system
type User struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	// HashedPassword []byte     `json:"-"`
	// PasswordType   string     `json:"-"`
	// AuthToken      string     `json:"auth_token"`
	// SystemAdmin    bool       `json:"system_admin"`
	// CreatedAt      time.Time  `json:"created_at"`
	// UpdatedAt      time.Time  `json:"updated_at"`
	// DeletedAt      *time.Time `json:"-"`
}

// Count is a count of the number of rows
type Count struct {
	Count int `json:"count"`
}
