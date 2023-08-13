package routes

import (
	"github.com/BrightOlawale/Inventory-API/handlers"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up all the routes for the application
func SetupRoutes(app *fiber.App) {
	// Get all items
	app.Get("/api/v1/items", handlers.GetAllItems)
	// Get item by ID
	app.Get("/api/v1/items/:id", handlers.GetItemByID)
	// Create item
	app.Post("api/v1/items", handlers.CreateItem)
	// Update item
	app.Put("api/v1/items/:id", handlers.UpdateItem)
	// Delete item
	app.Delete("api/v1/items/:id", handlers.DeleteItem)
}
