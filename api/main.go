package main

import (
    "os"
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

func main() {
    host     := os.Getenv("DB_HOST")
    port     := 5432
    user     := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname   := os.Getenv("DB_NAME")
        // connection string
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    fmt.Println(psqlconn)

        // open database
    db, err := sql.Open("postgres", psqlconn)
    CheckError(err)

        // close database
    defer db.Close()

        // check db
    err = db.Ping()
    CheckError(err)

    fmt.Println("Connected!")
}

func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}