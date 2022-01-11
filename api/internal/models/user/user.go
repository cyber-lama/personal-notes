package user

import (
	"database/sql"
	"github.com/cyber-lama/personal-notes/api/internal/exceptions/exception"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
)

type User struct {
	ID        uint            `json:"id"`
	Username  *sql.NullString `json:"username"`
	Password  string          `json:"password,omitempty"`
	Email     string          `json:"email,omitempty"`
	CreatedAt time.Time       `json:"created_at,omitempty"`
	UpdatedAt time.Time       `json:"updated_at,omitempty"`
}

func (u *User) Create(db *sqlx.DB, l *logrus.Logger) (*User, error) {
	//Substituting the Moscow time in the fields CreatedAt & UpdatedAt
	moscowTime, err := u.timeNow()
	if err != nil {
		return nil, err
	}
	u.CreatedAt = moscowTime
	u.UpdatedAt = moscowTime
	//Check if there is a user with such mail in db
	err = u.checkUniqueness(db, l)
	if err != nil {
		return nil, err
	}
	err = db.QueryRow(`
		INSERT INTO users (email, password, created_at, updated_at) VALUES ($1, $2, $3, $4) returning id
	`, u.Email, u.Password, u.CreatedAt, u.UpdatedAt).Scan(&u.ID)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) checkUniqueness(db *sqlx.DB, l *logrus.Logger) error {
	var result User
	moscowTime, err := u.timeNow()
	if err != nil {
		return err
	}
	result.CreatedAt = moscowTime
	result.UpdatedAt = moscowTime
	err = db.Get(&result, "SELECT id FROM users where email = $1", u.Email)
	switch err {
	case nil:
		m := exception.ExFields{"email": "Данный email уже используется"}
		return &m
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

func (u *User) timeNow() (time.Time, error) {
	moscow, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().In(moscow), nil
}
