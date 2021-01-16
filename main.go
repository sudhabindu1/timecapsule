package main

import (
	"log"
	"os"
	"timecapsule/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	_, ok := os.LookupEnv("AES_KEY")
	if !ok {
		log.Fatalln("private key not set")
	}
	app := fiber.New()

	app.Get("/", handlers.AppHandler)
	app.Post("/tc", handlers.CreateTimeCapsule)
	app.Get("/tc/:ciphertext", handlers.GetTimeCapsule)

	app.Listen("0.0.0.0:8080")
}
