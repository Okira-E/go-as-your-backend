package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func pingHandler(api fiber.Router) {
	routerPath := "/ping"

	api.Get(routerPath, ping)
}

func ping(c *fiber.Ctx) error {
	return c.Status(200).SendString("pong")
}