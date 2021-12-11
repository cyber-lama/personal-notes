package main

import (
	"github.com/cyber-lama/personal-notes/api/internal/apiserver"
	"github.com/joho/godotenv"
	"log"
)

//"github.com/golang-migrate/migrate"
//"github.com/golang-migrate/migrate/database/postgres"
//func upMigrate() error {
//	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
//		"db", "5432", "user", "test21", "postgres")
//	db, err := sql.Open("postgres", psqlconn)
//	if err != nil {
//		log.Fatal(err)
//		return err
//	}
//	driver, _ := postgres.WithInstance(db, &postgres.Config{})
//	m, err := migrate.NewWithDatabaseInstance(
//		"file://migrations",
//		"postgres", driver)
//	if err != nil {
//		log.Fatal(err)
//		return err
//	}
//	m.Up()
//	return nil
//}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	s := apiserver.New()

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
