package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kenwhite02/fiber-api/database"
	"github.com/kenwhite02/fiber-api/routes"
)

func setupRoutes(app *fiber.App) {
	// User endpoints
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":5000"))
}
