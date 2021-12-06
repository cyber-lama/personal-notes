package main

import (
	"github.com/cyber-lama/personal-notes/api/internal/apiserver"
	"log"
)

func main() {
	s := apiserver.New()

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
