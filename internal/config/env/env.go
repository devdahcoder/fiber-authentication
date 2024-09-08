package env

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

func EnvConfig(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	defer fmt.Printf("Env %s value has been returned", key)

	return os.Getenv(key)
}



