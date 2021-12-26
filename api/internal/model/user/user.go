package user

import "github.com/jmoiron/sqlx"

type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

func (r *User) Create(db *sqlx.DB) (*User, error) {
	return nil, nil
}
func (r *User) Find(db *sqlx.DB, id int) (*User, error) {
	return nil, nil
}

func (r *User) FindByEmail(db *sqlx.DB, email string) (*User, error) {
	return nil, nil
}
