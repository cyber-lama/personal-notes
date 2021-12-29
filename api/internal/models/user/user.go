package user

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type User struct {
	ID       uint            `json:"id"`
	Username *sql.NullString `json:"username"`
	Password string          `json:"password,omitempty"`
	Email    string          `json:"email,omitempty"`
}

func (u *User) Create(db *sqlx.DB, l *logrus.Logger) (*User, error) {
	_, err := u.checkUniqueness(db, l)
	if err != nil {
		return nil, err
	}
	err = db.QueryRow("INSERT INTO users (email, password) VALUES ($1, $2) returning id", u.Email, u.Password).Scan(&u.ID)

	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) checkUniqueness(db *sqlx.DB, l *logrus.Logger) (bool, error) {
	var result User
	err := db.Get(&result, "SELECT * FROM users where email = $1", u.Email)
	switch err {
	case nil:
		return false, errors.New("email: Данный email уже используется")
	case sql.ErrNoRows:
		return true, nil
	default:
		return false, err
	}
}

func (u *User) Find(db *sqlx.DB, id int) (*User, error) {
	return nil, nil
}

func (u *User) FindByEmail(db *sqlx.DB, email string) (*User, error) {
	return nil, nil
}
