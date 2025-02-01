package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// loads the .env file from the PROJECT_ROOT
func GetDotEnvValue(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		// use the logger from main.go, GetLogger()
		log.Fatalf("Error loading .env file")
		// log.GetLogger().Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}

// returns true if the ENV in .env is set to prod
func IsProd() bool {
	env := GetDotEnvValue("ENV")
	return env == "prod"
}
