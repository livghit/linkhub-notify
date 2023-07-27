package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/livghit/linkhub/pkg/config"
)

func Auth() func(c *fiber.Ctx) error {
	auth := basicauth.New(config.AuthConfig())

	return auth
}

func CSRF() func(c *fiber.Ctx) error {
	csrf := csrf.New(config.CsrfConfig())

	return csrf
}
