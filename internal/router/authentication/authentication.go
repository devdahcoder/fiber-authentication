package authentication

import (
	"github.com/gofiber/fiber/v2"
	"github/devdahcoder/fiber-authentication/internal/handler/authentication"
)

func AuthenticationRoute(app *fiber.App) {

	api := app.Group("/api/authentication")

	authenticationHandler := &authentication.AuthenticationStruct{}

	api.Post("/login", authenticationHandler.LoginHandler)

	api.Post("/signup", authenticationHandler.SignUpHandler)

}