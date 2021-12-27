package usercontroller

import (
	"github.com/jmoiron/sqlx"
)

type User struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) {

}
