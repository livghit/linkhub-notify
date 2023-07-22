package api

import (
	"github.com/gofiber/fiber/v2"
)

// In here we have all the stuff related to the api /api/v1/... all the routes and so on

var ApiRoutes map[string]func(c *fiber.Ctx);


func LoadApiRoutes(server *fiber.App) fiber.Router{
 // this will feed the Router with the API routes if existing 
  api := server.Group("api/v1");

  return api;
}

