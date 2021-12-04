package main

import (
    "github.com/cyber-lama/personal-notes/api/app/server"
)

func main() {
    err := server.UpServer()
    CheckError(err)
}

func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}