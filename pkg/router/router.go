package router

import "github.com/gofiber/fiber/v2"

/*
Kinda easy to understand im here I will create a simple router
*/

type Router struct {
	instance *fiber.App
}

func New() *fiber.App {
	r := Router{}
	r.instance = fiber.New()

	return r.instance
}
