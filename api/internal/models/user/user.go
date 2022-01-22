package user

import (
	"database/sql"
	"fmt"
	"github.com/cyber-lama/personal-notes/api/internal/exceptions/exception"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        uint            `json:"id"`
	Username  *sql.NullString `json:"username,omitempty"`
	Password  string          `json:"-"`
	Email     string          `json:"email"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Token     string          `json:"token"`
}

func (u *User) Create(db *sqlx.DB, l *logrus.Logger) (*User, error) {
	//Add the Moscow time in the fields CreatedAt & UpdatedAt
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
	// Create User
	err = db.QueryRow(`
		INSERT INTO users (email, password, created_at, updated_at) VALUES ($1, $2, $3, $4) returning id
	`, u.Email, "", u.CreatedAt, u.UpdatedAt).Scan(&u.ID)

	if err != nil {
		return nil, err
	}

	// create user token
	err = u.CreateToken(db, u.Email)
	if err != nil {
		return nil, err
	}

	// The user has already been saved,
	// using the goroutine we save the password asynchronously,
	// since hashing takes more than a second
	go u.hashPassword(db, u.Password)

	u.Sanitize()

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
	err = db.Get(&result, "SELECT id FROM users where email = $1 LIMIT 1", u.Email)
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

func (u *User) hashPassword(db *sqlx.DB, str string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(str), 14)
	db.QueryRow(`UPDATE users SET password = ($1) WHERE id = ($2)`, hash, u.ID)
}

func (u *User) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (u *User) CreateToken(db *sqlx.DB, email string) error {
	// hashing email, this is token
	b, _ := bcrypt.GenerateFromPassword([]byte(email), bcrypt.MinCost)
	moscowTime, _ := u.timeNow()
	strToken := fmt.Sprintf("%d|%s", u.ID, string(b))
	t := &Token{
		UserID:    u.ID,
		Token:     strToken,
		CreatedAt: moscowTime,
		UpdatedAt: moscowTime,
	}
	res, err := t.Create(db)
	if err != nil {
		return err
	}
	u.Token = string(res.Token)
	return nil
}

// Sanitize Clean Password property
func (u *User) Sanitize() {
	u.Password = ""
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
