package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func envMongoURI() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGOURI")
	return mongoURI
}
