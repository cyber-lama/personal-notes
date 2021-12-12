package store

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Store struct {
	config *DBConfig
	db     *sql.DB
}

func New(config *DBConfig) *Store {
	return &Store{config: config}
}

func (s *Store) Open() error {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		s.config.Host, s.config.Port, s.config.User, s.config.Password, s.config.DBname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *Store) Close() error {
	err := s.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) preparingDriver() (*migrate.Migrate, error) {
	driver, _ := postgres.WithInstance(s.db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (s *Store) UpMigrations() error {
	m, err := s.preparingDriver()
	err = m.Up()
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) DownMigrations() error {
	m, err := s.preparingDriver()
	err = m.Down()
	if err != nil {
		return err
	}
	return nil
}
