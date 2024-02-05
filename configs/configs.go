package configs

import (
	"os"

	"log"

	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func EnvMongoURI() string {
	return os.Getenv("MONGODB_URL")
}

func EnvPort() string {
	return os.Getenv("PORT")
}
