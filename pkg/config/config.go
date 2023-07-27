package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/livghit/linkhub/pkg/templating"
	"github.com/livghit/linkhub/pkg/utils"
)



type TemplatingEngine struct{
  ending string 
}


// Function to pass the Templating Engine to the fiber app
func ViewConfigs() fiber.Config{
  // maybe do engine instead of mustache and have all the engine under engine.NewDjango NewMustache etc?
  engine := templating.HtmlEngine()
  configs := fiber.Config{
    Views: engine,
    ViewsLayout: utils.Layout(),
  }

  return configs
}
