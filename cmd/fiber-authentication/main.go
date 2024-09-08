package main

import (
	"github/devdahcoder/fiber-authentication/internal/router/authentication"
	"github/devdahcoder/fiber-authentication/internal/config/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// initialize fiber framework
	app := fiber.New()
	// initialize database
	database.ConnectDB()
	// initialize authentication routers
	authentication.AuthenticationRoute(app)
	
	app.Listen(":3000")
}