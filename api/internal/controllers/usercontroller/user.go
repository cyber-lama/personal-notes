package usercontroller

import (
	"github.com/jmoiron/sqlx"
)

type UserController struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) {

}
