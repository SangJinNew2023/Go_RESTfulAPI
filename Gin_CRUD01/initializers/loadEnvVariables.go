package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// Dotenv load variables from a .env file into ENV when the environment is bootstrapped
func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
