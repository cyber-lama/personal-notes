package auth

import (
	"encoding/json"
	"github.com/cyber-lama/personal-notes/api/internal/exceptions/exception"
	"github.com/cyber-lama/personal-notes/api/internal/models/user"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
)

type AuthMiddleware struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

func New(r *http.Request) (*AuthMiddleware, error) {
	a := &AuthMiddleware{}
	if err := json.NewDecoder(r.Body).Decode(a); err != nil {
		m := exception.ExFields{"token": "Неверный формат данных"}
		return nil, m
	}
	return a, nil
}
func (a AuthMiddleware) tokenValidate(t *user.Token) error {
	return validation.ValidateStruct(
		t,
		validation.Field(
			&t.Token,
			validation.Required.Error("Поле token обязательно для заполнения"),
		),
	)
}
func (a AuthMiddleware) userIdValidate(c *user.User) error {
	return validation.ValidateStruct(
		c,
		validation.Field(
			&c.ID,
			validation.Required.Error("Поле id обязательно для заполнения"),
		),
	)
}

func (a AuthMiddleware) CheckAuth(db *sqlx.DB) error {
	// validate id of user
	u64, err := strconv.ParseUint(a.ID, 10, 64)
	if err != nil {
		m := exception.ExFields{"token": "Укажите корректный токен"}
		return m
	}
	c := &user.User{
		ID: uint(u64),
	}
	if err = a.userIdValidate(c); err != nil {
		return err
	}
	// validate token
	t := &user.Token{
		Token: []byte(a.Token),
	}
	if err = a.tokenValidate(t); err != nil {
		return err
	}
	//check token
	var token *user.Token
	token, err = t.CheckToken(db)
	if err != nil {
		return err
	}
	// ------------------------------------------
	if token.UserID != c.ID {
		m := exception.ExFields{"token": "Токен не соответствует пользователю"}
		return m
	}
	return nil
}
