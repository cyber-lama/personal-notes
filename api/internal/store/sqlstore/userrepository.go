package sqlstore

import "github.com/cyber-lama/personal-notes/api/internal/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	return nil
}
func (r *UserRepository) Find(id int) (*model.User, error) {
	return nil, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}
