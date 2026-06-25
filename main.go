package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/Tanu048/Admin_go/database"
	"github.com/Tanu048/Admin_go/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to the iSPARC API")
}

func setupRoutes(app *fiber.App) {
	// Basic test route
	app.Get("/api", welcome)

	// Authentication Routes
	auth := app.Group("/api/auth")
	auth.Post("/admin/login", routes.LoginAdmin)
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	setupRoutes(app)
	log.Println("Server starting on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
