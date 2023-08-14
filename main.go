package main

import (
	"fmt"
	"github.com/BrightOlawale/Inventory-API/database"
	"github.com/BrightOlawale/Inventory-API/routes"
	"github.com/BrightOlawale/Inventory-API/utils"
	"github.com/gofiber/fiber/v2"
)

// DEFAULT_PORT : default port to run the application on
const DEFAULT_PORT string = "500"

// NewFiberApp : function to create a new fiber application
func NewFiberApp() *fiber.App {
	// Create a new fiber application
	var app *fiber.App = fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	// Return the fiber application
	return app
}

// Main function: Entry point of the program
func main() {
	// Creating a fiber application
	var app *fiber.App = NewFiberApp()

	// Connect to the database
	database.ConnectDatabase(utils.GetValue("DB_NAME"))

	// Get application port from environment variable
	var port string = utils.GetValue("PORT")

	// If no port was set in the environment variable, then use the default port
	if port == "" {
		port = DEFAULT_PORT
	}

	// Starting the server at port 8080
	err := app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		return
	}
}
