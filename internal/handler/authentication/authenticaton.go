package authentication

import "github.com/gofiber/fiber/v2"

type AuthenticationInterface interface {
	loginHandler(c *fiber.Ctx) error
	signUpHandler(c *fiber.Ctx) error
}

type AuthenticationStruct struct {}

func (a *AuthenticationStruct) LoginHandler(c *fiber.Ctx) error {
	return c.SendString("Login page")
}

func (a *AuthenticationStruct) SignUpHandler(c *fiber.Ctx) error {
	return c.SendString("Sign up page")
}
