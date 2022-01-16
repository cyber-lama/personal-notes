package user

import "time"

type UserToken struct {
	ID        string
	UserID    string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *UserToken) HashToken() {

}
