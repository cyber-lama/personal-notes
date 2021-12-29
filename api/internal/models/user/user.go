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

//TODO not working
func (u *User) checkUniqueness(db *sqlx.DB, l *logrus.Logger) (bool, error) {
	e := sql.Row{}
	q := db.QueryRow("SELECT * FROM users where email = $1", u.Email)
	fmt.Println("empty row", e)
	fmt.Println("select", q)
	if &e == q {
		return true, nil
	}

	return false, errors.New("email: Данный email уже используется")
}

func (u *User) Find(db *sqlx.DB, id int) (*User, error) {
	return nil, nil
}

func (u *User) FindByEmail(db *sqlx.DB, email string) (*User, error) {
	return nil, nil
}
