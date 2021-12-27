package controllers

import (
	"encoding/json"
	"net/http"
)

type BaseController struct {
}

func New() *BaseController {
	return &BaseController{}
}

func (c BaseController) Respond(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			return
		}
	}
}

func (c BaseController) Error(w http.ResponseWriter, code int, err error) {
	c.Respond(w, code, map[string]interface{}{"status": code, "errors": err})
}
