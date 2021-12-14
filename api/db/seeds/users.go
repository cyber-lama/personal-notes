package seeds

import (
	"github.com/bxcodec/faker/v3"
	"log"
)

func (s Seed) UsersSeed() {
	for i := 0; i < 50; i++ {
		_, err := s.db.Exec(`
			INSERT INTO users(username, password, email) VALUES ($1, $2, $3)`,
			faker.FirstName(), faker.Password(), faker.Email())
		if err != nil {
			log.Fatalf("error seeding roles: %v", err)
		}
	}
}
