package store

import (
	"github.com/jmoiron/sqlx"
)

type Store struct {
	DB *sqlx.DB
}

// New init store struct with db connect
func New(db *sqlx.DB) *Store {
	return &Store{
		DB: db,
	}
}
