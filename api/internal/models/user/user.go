package user

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

func (u *User) Create(db *sqlx.DB, l *logrus.Logger) (*User, error) {
	query, err := db.Query("INSERT INTO users (username, password, Email) VALUES ($1, $2, $3)", u.Username, u.Password, u.Email)
	if err != nil {
		l.Error("db.Query error ", err)
		return nil, err
	}
	fmt.Println(query.NextResultSet())
	return u, nil
}
func (u *User) Find(db *sqlx.DB, id int) (*User, error) {
	return nil, nil
}

func (u *User) FindByEmail(db *sqlx.DB, email string) (*User, error) {
	return nil, nil
}
