package handlers

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"timecapsule/model"
	"timecapsule/service"

	"github.com/gofiber/fiber/v2"
)

func CreateTimeCapsule(c *fiber.Ctx) error {
	m := new(model.Message)
	err := c.BodyParser(&m)
	if err != nil {
		return fiber.ErrBadRequest
	}

	if !m.IsValid() {
		return fiber.ErrBadRequest
	}

	cipherText, err := service.Encrypt(fmt.Sprintf("%s-%v-%s", m.Sender, m.RevealMillis, m.Body))
	if err != nil {
		return fiber.ErrInternalServerError
	}

	baseUrl := "http://localhost:8080"
	_, ok := os.LookupEnv("PORT")
	if ok {
		baseUrl = "https://timecapsul.herokuapp.com"
	}

	return c.SendString(fmt.Sprintf("%s/tc/%s", baseUrl, cipherText))
}

func GetTimeCapsule(c *fiber.Ctx) error {
	cipherText := c.Params("ciphertext")
	plainText, err := service.Decrypt(cipherText)
	if err != nil {
		return fiber.ErrBadRequest
	}

	plainTextArr := strings.SplitN(plainText, "-", 3)
	if len(plainTextArr) != 3 {
		return fiber.ErrBadRequest
	}

	revealTs, err := strconv.ParseInt(plainTextArr[1], 10, 64)
	if err != nil {
		return fiber.ErrBadRequest
	}

	if revealTs > time.Now().UnixNano()/1000000 {
		return c.SendString(fmt.Sprintf("Time capsule is sealed till: %v", revealTs))
	}

	m := model.Message{
		Sender: plainTextArr[0],
		Body:   plainTextArr[2],
	}

	return c.JSON(m)
}
