package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/livghit/linkhub/pkg/config"
	"github.com/livghit/linkhub/pkg/router"
)

/*

this is the place where we define our server
in here we will pull the configs and render all the necesary env in order to connect to the database , check if production or not (this should go inside utils I guess but we will see )

this will have a struct called Server and a method (constructor) to spin up a Server ( using the Fiber app )

*/

// the  server will have an router , server configs[link to litsen too , app name and so on]
type Server struct {
	router  *fiber.App
	configs config.ServerConfigs
}

func New() *Server {
	s := Server{
		router:  router.New(),
		configs: config.ServerConfigs{},
	}

	return &s
}
