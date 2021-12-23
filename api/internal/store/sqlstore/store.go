package sqlstore

import (
	"github.com/cyber-lama/personal-notes/api/internal/store"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db             *sqlx.DB
	userRepository *UserRepository
}

// New init store struct with db connect and all repo
func New(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

// User add to Store property userRepository with UserRepository struct
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
