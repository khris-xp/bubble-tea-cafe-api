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
	loadEnv()
	return os.Getenv("MONGODB_URL")
}

func EnvPort() string {
	loadEnv()
	return os.Getenv("PORT")
}

func EnvSecretKey() string {
	loadEnv()
	return os.Getenv("JWT_SECRET")
}
