package get_caption

import (
	"github.com/gofiber/fiber/v2"
)

func HttpGetCaption() {
	response := fiber.Get("https://api.imgflip.com/get_memes")
}
