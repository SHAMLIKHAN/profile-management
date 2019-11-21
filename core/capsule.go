package core

import (
	"database/sql"
)

// Capsule : Capsule of the app, all functions are signed with this struct
type Capsule struct {
	DB *sql.DB
}

// GetCapsule : Just to return Capsule struct
func GetCapsule(db *sql.DB) *Capsule {
	return &Capsule{
		DB: db,
	}
}
