package user

import (
	"database/sql"
	"github.com/cyber-lama/personal-notes/api/internal/exceptions/exception"
	"github.com/jmoiron/sqlx"
	"time"
)

type Token struct {
	ID        uint      `db:"id"`
	UserID    uint      `db:"user_id"`
	Token     string    `db:"token"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (t *Token) Create(db *sqlx.DB) (*Token, error) {
	err := db.QueryRow(`
		INSERT INTO users_token (user_id, token, created_at, updated_at) VALUES ($1, $2, $3, $4) returning id
	`, t.UserID, t.Token, t.CreatedAt, t.UpdatedAt).Scan(&t.ID)

	if err != nil {
		return nil, err
	}
	return t, nil
}

func (t *Token) CheckToken(db *sqlx.DB) (*Token, error) {
	err := db.Get(t, "SELECT * FROM users_token where token = $1 LIMIT 1", t.Token)
	switch err {
	case nil:
		// Token there is in db
		return t, nil
	case sql.ErrNoRows:
		// No token in db
		m := exception.ExFields{"token": "Токен не существует"}
		return nil, m
	default:
		// Error
		return nil, err
	}
}
