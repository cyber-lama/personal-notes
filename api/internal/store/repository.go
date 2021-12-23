package store

import "github.com/cyber-lama/personal-notes/api/internal/model"

// UserRepository interface for work with User model
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
