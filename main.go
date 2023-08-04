package main

import (
	"github.com/gofiber/fiber/v2"
)

// Main function: Entry point of the program
func main() {
	// Creating a fiber application
	var app *fiber.App = fiber.New()

	// Adding a request handler for the "/" route
	app.Get("/", func(ctx *fiber.Ctx) error {
		// Returning "Welcome to Fiber!" as the response
		return ctx.SendString("Welcome to Fiber!")
	})

	// Starting the server at port 8080
	err := app.Listen(":8080")
	if err != nil {
		return
	}
}
