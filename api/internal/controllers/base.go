package controllers

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type BaseController struct {
	logger *logrus.Logger
}

func New(l *logrus.Logger) *BaseController {
	return &BaseController{
		logger: l,
	}
}

func (c BaseController) Respond(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			return
		}
	}
}
func (c BaseController) Message(w http.ResponseWriter, code int, data interface{}) {
	c.Respond(w, code, map[string]interface{}{"status": code, "data": data})
}
func (c BaseController) Error(w http.ResponseWriter, code int, err error) {
	c.Respond(w, code, map[string]interface{}{"status": code, "errors": err})
}
