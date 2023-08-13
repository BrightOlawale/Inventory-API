package main

import (
	"github.com/BrightOlawale/Inventory-API/routes"
	"github.com/gofiber/fiber/v2"
)

// Main function: Entry point of the program
func main() {
	// Creating a fiber application
	var app *fiber.App = fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	// Starting the server at port 8080
	err := app.Listen(":8080")
	if err != nil {
		return
	}
}
