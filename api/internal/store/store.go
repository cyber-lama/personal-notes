package store

import (
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

// New init store struct with db connect
func New(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}
