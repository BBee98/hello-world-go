package routes

import (
	"github.com/gofiber/fiber/v2"
	"hello-world/backend/config"
)

func Home() {
	config.App.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})
}
