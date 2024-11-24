package main

import (
	"inventory-management-system/config"
	"inventory-management-system/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDatabase()

	app := fiber.New()

	routes.RegisterRoutes(app)

	log.Println("Server is running on http://127.0.0.1:4000")
	log.Fatal(app.Listen(":4000"))
}
