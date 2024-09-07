package main

import (
	"fmt"
	"github/devdahcoder/fiber-authentication/internal/router/authentication"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"os"
)

func Config(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	app := fiber.New()

	authentication.AuthenticationRoute(app)

	app.Listen(":3000")
}