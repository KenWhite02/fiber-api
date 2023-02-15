package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kenwhite02/fiber-api/database"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to Fiber API.")
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/api", welcome)

	log.Fatal(app.Listen(":5000"))
}
