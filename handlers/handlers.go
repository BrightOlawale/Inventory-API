package handlers

import (
	"github.com/BrightOlawale/Inventory-API/models"
	"github.com/BrightOlawale/Inventory-API/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// GetAllItems is the handler for the GET /items route
func GetAllItems(c *fiber.Ctx) error {
	// Get all Item from storage
	var items []models.Item = service.GetAllItems()

	// Create a response object
	// The response type is Response[[]models.Item]
	var response models.Response[[]models.Item] = models.Response[[]models.Item]{
		Success: true,
		Message: "All items",
		Data:    items,
	}

	// Return the response as JSON
	return c.Status(http.StatusOK).JSON(response)
}

// GetItemByID : is the handler for the GET /items/:id route
func GetItemByID(c *fiber.Ctx) error {
	// Get ID from the request parameters
	var itemId string = c.Params("id")

	// Use the GetItemById from the service package to get the item from storage
	item, err := service.GetItemById(itemId)

	// If no error occurred
	if err != nil {
		// Create response object
		var response models.Response[any] = models.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}

		// Now set status code and return response
		return c.Status(http.StatusNotFound).JSON(response)
	}

	// if everything was fine then create response object
	var response models.Response[models.Item] = models.Response[models.Item]{
		Success: true,
		Message: "Item found",
		Data:    item,
	}

	// We return the response and indicate response was OK 200
	return c.Status(http.StatusOK).JSON(response)
}

// CreateItem : is the handler for the POST /items route
func CreateItem(c *fiber.Ctx) error {
	// Create variable to hold request body
	var itemRequest *models.ItemRequest = new(models.ItemRequest)

	// Parse body into itemRequest
	// What BodyParser does is to parse the request body into the itemRequest variable
	// If there is an error, it will return the error
	if err := c.BodyParser(itemRequest); err != nil {
		// Create response object
		var response models.Response[any] = models.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}

		// Return response and indicate bad request
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// If BodyParser was successful, then validate the itemRequest
	// ValidateStruct is a method we created in the ItemRequest model
	// It returns an array of ErrorResponse
	var validationErrors []*models.ErrorResponse = itemRequest.ValidateStruct()

	// If there is an error
	if len(validationErrors) > 0 {
		// Create response object
		var response models.Response[any] = models.Response[any]{
			Success: false,
			Message: "Validation error",
			Data:    validationErrors,
		}

		// Return response and indicate bad request
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// If there is no error, then create the item
	var newItem models.Item = service.CreateItem(*itemRequest)

	// Create response object
	var response models.Response[models.Item] = models.Response[models.Item]{
		Success: true,
		Message: "Item created",
		Data:    newItem,
	}

	// Return response and indicate item was created
	return c.Status(http.StatusCreated).JSON(response)
}

// UpdateItem : is the handler for the PUT /items/:id route
func UpdateItem(c *fiber.Ctx) error {
	// I think it makes sense to first sanity check the ID
	// Get ID from the request parameters
	var itemId string = c.Params("id")

	// Use the GetItemById from the service package to get the item from storage
	_, err := service.GetItemById(itemId)

	// If we don't find the item, then that means the ID is invalid
	// So we return a 404 Not Found
	if err != nil {
		// Create response object
		var response models.Response[any] = models.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}

		// Now set status code and return response
		return c.Status(http.StatusNotFound).JSON(response)
	}

	// Now that we know the ID is valid, we can proceed to update the item
	// Create variable to hold request body
	var itemRequest *models.ItemRequest = new(models.ItemRequest)

	// Parse body into itemRequest using BodyParser
	if err := c.BodyParser(itemRequest); err != nil {
		// Create response object
		var response models.Response[any] = models.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}

		// Return response and indicate bad request
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// If BodyParser was successful, then validate the itemRequest using ValidateStruct
	var validationErrors []*models.ErrorResponse = itemRequest.ValidateStruct()

	// If there is an error
	if len(validationErrors) > 0 {
		// Create response object
		var response models.Response[any] = models.Response[any]{
			Success: false,
			Message: "Validation error",
			Data:    validationErrors,
		}

		// Return response and indicate bad request
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// If there is no error, then update the item
	updatedItem, err := service.UpdateItem(*itemRequest, itemId)

	// If there is an error
	if err != nil {
		// Create response object
		var response models.Response[any] = models.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}

		// Return response and indicate bad request
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// Create response object
	var response models.Response[models.Item] = models.Response[models.Item]{
		Success: true,
		Message: "Item updated",
		Data:    updatedItem,
	}

	// Return response and indicate item was updated
	return c.Status(http.StatusOK).JSON(response)
}

// DeleteItem : is the handler for the DELETE /items/:id route
func DeleteItem(c *fiber.Ctx) error {
	// Get ID from the request parameters
	var itemId string = c.Params("id")

	// Use the GetItemById from the service package to get the item from storage
	_, err := service.GetItemById(itemId)

	// If we don't find the item, then that means the ID is invalid
	// So we return a 404 Not Found
	if err != nil {
		// Create response object
		var response models.Response[any] = models.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}

		// Now set status code and return response
		return c.Status(http.StatusNotFound).JSON(response)
	}

	// If there is no error, then delete the item
	var deleteResponse bool = service.DeleteItem(itemId)

	// If deleteResponse is false, then that means the item was not deleted
	// So we return a 500 Internal Server Error
	if !deleteResponse {
		// Create response object
		var response models.Response[any] = models.Response[any]{
			Success: false,
			Message: "Internal server error",
			Data:    nil,
		}

		// Now set status code and return response
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	// But if deleteResponse is true, then that means the item was deleted
	// So we return a 200 OK
	// Create response object
	var response models.Response[any] = models.Response[any]{
		Success: true,
		Message: "Item deleted",
		Data:    nil,
	}

	// Now set status code and return response
	return c.Status(http.StatusOK).JSON(response)
}
