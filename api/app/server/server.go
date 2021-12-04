package server

import (
    "fmt"
    "os"
    "net/http"
    "github.com/cyber-lama/personal-notes/api/app/routes"
)

func UpServer () {
    httpPort := os.Getenv("PORT")
    http.ListenAndServe(":" + httpPort, routes.HandleReq)
    fmt.Println("Server listening!")
}