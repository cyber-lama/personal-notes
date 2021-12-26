package authcontroller

import (
	"encoding/json"
	"github.com/cyber-lama/personal-notes/api/internal/store"
	"net/http"
)

type AuthController struct {
	db *store.Store
}

func New(db *store.Store) *AuthController {
	return &AuthController{
		db: db,
	}
}

type Dog struct {
	Name string
}

func (c AuthController) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		d := Dog{
			Name: "Borky Borkins",
		}

		c.respond(w, r, http.StatusOK, d)
	}
}
func (c AuthController) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			return
		}
	}
}
