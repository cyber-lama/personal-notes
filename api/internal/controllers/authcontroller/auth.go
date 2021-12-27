package authcontroller

import (
	"encoding/json"
	"github.com/cyber-lama/personal-notes/api/internal/controllers"
	"github.com/cyber-lama/personal-notes/api/internal/models/user"
	"github.com/cyber-lama/personal-notes/api/internal/store"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/sirupsen/logrus"
	"net/http"
)

type AuthController struct {
	base   *controllers.BaseController
	db     *store.Store
	logger *logrus.Logger
}

func New(db *store.Store, l *logrus.Logger) *AuthController {
	b := controllers.New()
	return &AuthController{
		base:   b,
		db:     db,
		logger: l,
	}
}

func (c AuthController) registerValidate(u *user.User) error {
	return validation.ValidateStruct(
		u,
		validation.Field(
			&u.Email,
			validation.Required.Error("Поле email обязательно для заполнения"),
			is.Email.Error("Поле должно быть валидным email"),
			validation.Length(4, 100).Error("Длинна поля email от 4 до 100 символов"),
		),
		validation.Field(
			&u.Username,
			validation.Required.Error("Поле username обязательно для заполнения"),
			validation.Length(2, 100).Error("Длинна поля username от 2 до 100 символов"),
		),
		validation.Field(
			&u.Password,
			validation.Required.Error("Поле password обязательно для заполнения"),
			validation.Length(6, 100).Error("Длинна поля password от 2 до 100 символов"),
		),
	)
}

func (c AuthController) Register() http.HandlerFunc {
	type request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			c.base.Error(w, http.StatusBadRequest, err)
			return
		}
		u := &user.User{
			Email:    req.Email,
			Password: req.Password,
			Username: req.Username,
		}
		if err := c.registerValidate(u); err != nil {
			c.base.Error(w, http.StatusBadRequest, err)
			return
		}
		res, err := u.Create(c.db.DB, c.logger)
		if err != nil {
			c.base.Error(w, http.StatusBadRequest, err)
			return
		}
		c.base.Respond(w, http.StatusOK, res)
	}
}
