package server

import (
	"github.com/gofiber/fiber/v2"
)

func StartServer() {

	app := fiber.New()
	app.Get("/download", DownloadHandler)
	app.Listen(":3000")
}
