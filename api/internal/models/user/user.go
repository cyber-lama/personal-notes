package user

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
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
	e := sql.Row{}
	q := db.QueryRow("SELECT * FROM users where email = $1", u.Email)
	if &e == q {
		return true, nil
	}
	err := errors.New("email: Данный email уже используется")
	fmt.Print(err)
	return false, err
}

func (u *User) Find(db *sqlx.DB, id int) (*User, error) {
	return nil, nil
}

func (u *User) FindByEmail(db *sqlx.DB, email string) (*User, error) {
	return nil, nil
}
