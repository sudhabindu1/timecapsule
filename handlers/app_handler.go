package handlers

import "github.com/gofiber/fiber/v2"

func AppHandler(c *fiber.Ctx) error {
	return c.SendString("Time Capsule ")
}
