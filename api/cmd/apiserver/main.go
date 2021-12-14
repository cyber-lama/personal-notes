package main

import (
	"github.com/cyber-lama/personal-notes/api/internal/apiserver"
	"github.com/joho/godotenv"
	"log"
)

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
