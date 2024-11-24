package main

import (
	"github.com/shohan-pherones/inventory-management-system.git/config"
	"github.com/shohan-pherones/inventory-management-system.git/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	config.ConnectDatabse()

	routes.RegisterRoutes(app)

	app.Listen(":4000")
}
