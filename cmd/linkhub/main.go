package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/livghit/linkhub/pkg/config"
	"github.com/livghit/linkhub/web/handlers"
	"log"
)

/*
QUESTIONS :
- Where will the middleware will be placed
- How to handle user login with azur ad
*/

func main() {

	app := fiber.New(config.ViewConfigs())

	// From here register all routes
	app.Get("/", handlers.HomepageHandler)

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.Render("pages/index", fiber.Map{
			"Title": "App",
		}, "layouts/base")
	})

	app.Get("/users", handlers.UserHandler)

	log.Fatal(app.Listen(":3000"))

}
