package apiserver

import (
	"github.com/cyber-lama/personal-notes/api/internal/store"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func Start(c *Config) error {
	db, err := newDB(c.DatabaseURL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
	st := store.New(db)
	srv := newServer(st)

	return http.ListenAndServe(c.HTTPPort, srv)
}

func newDB(dbURL string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db, nil
}
