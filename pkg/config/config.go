package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/livghit/linkhub/pkg/templating"
	"github.com/livghit/linkhub/web/handlers"
)

type TemplatingEngine struct {
	ending string
}

// Function to pass the Templating Engine to the fiber app
func ViewConfigs() fiber.Config {
	// maybe do engine instead of mustache and have all the engine under engine.NewDjango NewMustache etc?
	engine := templating.HtmlEngine()
  engine.Reload(true)
	configs := fiber.Config{
		Views:       engine,
	}

	return configs
}

func AuthConfig() basicauth.Config {
	authconfig := basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "123456",
		},
    Unauthorized: handlers.NotAuthorized,
	}

	return authconfig
}
