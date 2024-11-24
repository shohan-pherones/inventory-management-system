package routes

import (
	"inventory-management-system/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	productGroup := app.Group("/products")

	productGroup.Post("/", controllers.CreateProduct)
}
