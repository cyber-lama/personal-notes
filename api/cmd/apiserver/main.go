package main

import (
	"fmt"
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

//func init() {
//	err := godotenv.Load()
//	if err != nil {
//		log.Fatal(err)
//	}
//}

func main() {
	fmt.Println("test")
	//s := apiserver.New()
	//
	//if err := s.Start(); err != nil {
	//	log.Fatal(err)
	//}
}
