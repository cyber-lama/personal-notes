package main

import (
    "github.com/cyber-lama/personal-notes/api/app/routes"
)

func main() {
    // start listen req
    routes.HandleReq()
}

func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}