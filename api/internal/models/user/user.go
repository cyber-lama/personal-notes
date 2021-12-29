package user

import (
	"database/sql"
	"encoding/json"
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
	err := u.checkUniqueness(db, l)
	if err != nil {
		return nil, err
	}
	err = db.QueryRow("INSERT INTO users (email, password) VALUES ($1, $2) returning id", u.Email, u.Password).Scan(&u.ID)

	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) checkUniqueness(db *sqlx.DB, l *logrus.Logger) error {
	var result User
	err := db.Get(&result, "SELECT * FROM users where email = $1", u.Email)
	switch err {
	case nil:
		errStr := map[string]string{"email": "Данный email уже используется"}
		errVar, _ := json.Marshal(errStr)
		l.Error(string(errVar))
		return errors.New(string("email: Данный email уже используется"))
	case sql.ErrNoRows:
		return nil
	default:
		return err
	}
}

func (u *User) Find(db *sqlx.DB, id int) (*User, error) {
	return nil, nil
}

func (u *User) FindByEmail(db *sqlx.DB, email string) (*User, error) {
	return nil, nil
}
