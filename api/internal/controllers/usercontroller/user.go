package usercontroller

import (
	"github.com/cyber-lama/personal-notes/api/internal/controllers"
	"github.com/cyber-lama/personal-notes/api/internal/middleware/auth"
	"github.com/cyber-lama/personal-notes/api/internal/store"
	"github.com/sirupsen/logrus"
	"net/http"
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

func (u UserController) GetInfo() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		middleware, err := auth.New(r)
		if err != nil {
			u.base.Error(w, http.StatusBadRequest, err)
			return
		}
		err = middleware.CheckAuth(u.db.DB)
		if err != nil {
			u.base.Error(w, http.StatusBadRequest, err)
			return
		}
		u.logger.Infoln("Успешный запрос")
	}
}
