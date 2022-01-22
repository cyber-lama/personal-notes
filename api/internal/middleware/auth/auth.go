package auth

import (
	"encoding/json"
	"github.com/cyber-lama/personal-notes/api/internal/exceptions/exception"
	"github.com/cyber-lama/personal-notes/api/internal/models/user"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jmoiron/sqlx"
	"net/http"
	"regexp"
)

type AuthMiddleware struct {
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
	reToken := regexp.MustCompile("^(\\d*)(\\|)(.*)$")
	return validation.ValidateStruct(
		t,
		validation.Field(
			&t.Token,
			validation.Match(reToken).Error("Некорректный токен"),
			validation.Required.Error("Поле token обязательно для заполнения"),
		),
	)
}

func (a AuthMiddleware) CheckAuth(db *sqlx.DB) error {

	// validate token
	t := &user.Token{
		Token: a.Token,
	}
	if err := a.tokenValidate(t); err != nil {
		return err
	}
	//check token
	_, err := t.CheckToken(db)
	if err != nil {
		return err
	}
	// ------------------------------------------
	//if token.UserID != c.ID {
	//	m := exception.ExFields{"token": "Токен не соответствует пользователю"}
	//	return m
	//}
	return nil
}
