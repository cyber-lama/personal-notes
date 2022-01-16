package user

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type Token struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Token     []byte    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
