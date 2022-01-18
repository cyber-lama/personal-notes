package usercontroller

import (
	"encoding/json"
	"github.com/cyber-lama/personal-notes/api/internal/controllers"
	"github.com/cyber-lama/personal-notes/api/internal/exceptions/exception"
	"github.com/cyber-lama/personal-notes/api/internal/models/user"
	"github.com/cyber-lama/personal-notes/api/internal/store"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type UserController struct {
	base   *controllers.BaseController
	db     *store.Store
	logger *logrus.Logger
}

func New(db *store.Store, l *logrus.Logger) *UserController {
	b := controllers.New(l)
	return &UserController{
		base:   b,
		db:     db,
		logger: l,
	}
}
func (u UserController) tokenValidate(t *user.Token) error {
	return validation.ValidateStruct(
		t,
		validation.Field(
			&t.Token,
			validation.Required.Error("Поле token обязательно для заполнения"),
		),
	)
}

func (u UserController) userIdValidate(c *user.User) error {
	return validation.ValidateStruct(
		c,
		validation.Field(
			&c.ID,
			validation.Required.Error("Поле id обязательно для заполнения"),
		),
	)
}

func (u UserController) GetInfo() http.HandlerFunc {
	type request struct {
		ID    string `json:"id"`
		Token string `json:"token"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			u.base.Error(w, http.StatusBadRequest, err)
			return
		}
		// validate id of user
		u64, err := strconv.ParseUint(req.ID, 10, 64)
		if err != nil {
			m := exception.ExFields{"id": "Введите корректный id пользователя"}
			u.base.Error(w, http.StatusBadRequest, m)
			return
		}
		c := &user.User{
			ID: uint(u64),
		}
		if err = u.userIdValidate(c); err != nil {
			u.base.Error(w, http.StatusBadRequest, err)
			return
		}
		// validate token
		t := &user.Token{
			Token: []byte(req.Token),
		}
		if err = u.tokenValidate(t); err != nil {
			u.base.Error(w, http.StatusBadRequest, err)
			return
		}

		//check token
		var token *user.Token
		token, err = t.CheckToken(u.db.DB)
		if err != nil {
			u.base.Error(w, http.StatusBadRequest, err)
			return
		}
		if token.UserID != c.ID {
			m := exception.ExFields{"token": "Токен не соответствует пользователю"}
			u.base.Error(w, http.StatusBadRequest, m)
			return
		}
	}
}
