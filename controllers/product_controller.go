package controllers

import (
	"context"
	"inventory-management-system/config"
	"inventory-management-system/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if product.Name == "" || product.Quantity <= 0 || product.Price <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input. Ensure Name, Quantity, and Price are valid.",
		})
	}

	product.ID = primitive.NewObjectID()

	collection := config.DB.Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create product",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}

func GetAllProducts(c *fiber.Ctx) error {
	collection := config.DB.Collection("products")

	var products []models.Product

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve products",
		})
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to parse product data",
			})
		}
		products = append(products, product)
	}

	if len(products) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No products found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(products)
}

func GetProductByID(c *fiber.Ctx) error {
	productID := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID format",
		})
	}

	collection := config.DB.Collection("products")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var product models.Product
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&product)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	productID := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID format",
		})
	}

	var updatedProduct models.Product

	if err := c.BodyParser(&updatedProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to parse request body",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": objectID}

	update := bson.M{
		"$set": bson.M{
			"name":     updatedProduct.Name,
			"quantity": updatedProduct.Quantity,
			"price":    updatedProduct.Price,
		},
	}

	collection := config.DB.Collection("products")

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update product",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product updated successfully",
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	productID := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID format",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := config.DB.Collection("products")

	filter := bson.M{"_id": objectID}

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete product",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}
