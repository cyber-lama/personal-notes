package authcontroller

import (
	"encoding/json"
	"github.com/cyber-lama/personal-notes/api/internal/models/user"
	"github.com/cyber-lama/personal-notes/api/internal/store"
	"github.com/sirupsen/logrus"
	"net/http"
)

type AuthController struct {
	db     *store.Store
	logger *logrus.Logger
}

func New(db *store.Store, l *logrus.Logger) *AuthController {
	return &AuthController{
		db:     db,
		logger: l,
	}
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
			c.logger.Error("r.Body decode error ", err)
			c.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &user.User{
			Email:    req.Email,
			Password: req.Password,
			Username: req.Username,
		}
		res, err := u.Create(c.db.DB, c.logger)
		if err != nil {
			c.logger.Error("user create error ", err)
			return
		}
		c.respond(w, r, http.StatusOK, res)
	}
}
func (c AuthController) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			c.logger.Error("res encode error ", err)
			return
		}
	}
}

func (c AuthController) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	c.respond(w, r, code, map[string]string{"error": err.Error()})
}
