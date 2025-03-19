package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

var App *fiber.App

func InitConfig() {

	engine := html.New("./views", ".html")

	App = fiber.New(fiber.Config{
		Views: engine,
		// Default global path to search for views (can be overriden when calling Render())
	})

}
