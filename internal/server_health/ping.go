package server_health

import (
	"github.com/gofiber/fiber/v2"
	"github.com/org/example/internal/utils"
)

func SetupHandlers(api fiber.Router) {
	routerPath := "/ping"

	api.Get(routerPath, ping)
}

func ping(c *fiber.Ctx) error {
	return utils.Ok(c, 200, "PONG", nil)
}