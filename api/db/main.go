package main

import (
	"fmt"
	"github.com/cyber-lama/personal-notes/api/db/seeds"
	"github.com/danvergara/seeder"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

//TODO необходимо вынести подключение к бд в отдельный класс
func main() {
	// connect to db
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sqlx.Connect("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	s := seeds.NewSeed(db)

	//run seeds
	if err := seeder.Execute(s); err != nil {
		log.Fatalf("error seeding the db %s\n", err)
	}
}
