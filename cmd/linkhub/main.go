package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/livghit/linkhub/web/templates"
)

/*

This will buill the app maybe I will a a bootstrap under the  utils.go Ill have to see

This will look something like that

router := router.New();
serverConfigs := configs.LoadConfigs()
server := server.New(serverConfigs , router)

server.Run();


QUESTIONS :
- Where will the middleware will be placed
- How to handle user login with azur ad
*/

func main() {

	app := fiber.New()

	app.Get("/homepage", web.HomepageHandler)

	log.Fatal(app.Listen(":3000"))
}
