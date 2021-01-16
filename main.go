package main

import (
	"os"
	"timecapsule/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	os.Setenv("AES_KEY", "64a46629c79692eaa20828ce6a1d90fa5e7a1011a68e18128aea8fddb7a5e018")
	app := fiber.New()

	app.Get("/", handlers.AppHandler)
	app.Post("/tc", handlers.CreateTimeCapsule)
	app.Get("/tc/:ciphertext", handlers.GetTimeCapsule)

	app.Listen(":8080")
}
