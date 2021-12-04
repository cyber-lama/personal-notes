package routes

import (
    "os"
    "net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type User struct {
    ID        int64  `json:"id"`
    Email     string `json:"email"`
    FirstName string `json:"first_name"`
}

func testRes (w http.ResponseWriter, r *http.Request) {
        u := User{}
        u.ID = 22
        u.Email = "test@test.com"
        u.FirstName = "John"

        us, _ := json.Marshal(u)

        w.Header().Set("Content-Type", "application/json")
        w.Write(us)
}

func HandleReq () {
    // port from docker-compose environment
    httpPort := os.Getenv("PORT")

    r := mux.NewRouter().StrictSlash(true)
    r.HandleFunc("/test", testRes).Methods("GET")

    http.ListenAndServe(":" + httpPort, r)
}