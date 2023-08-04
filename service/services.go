package service

import (
	"errors"
	"github.com/BrightOlawale/Inventory-API/models"
	"github.com/google/uuid"
	"time"
)

// Create variable to store item data
var storage []models.Item = []models.Item{}

// GetAllItems : get all items from the storage.
func GetAllItems() []models.Item {
	return storage
}

// GetItemById : get the item’s data by ID.
func GetItemById(id string) (models.Item, error) {
	// We iterate thru all the items in storage
	for _, item := range storage {
		// If the current item's ID is exactly same as the id parameter
		if id == item.ID {
			// Return item's data
			return item, nil
		}
	}
	// If item was not found return error
	return models.Item{}, errors.New("item was not found")
}

// CreateItem : function to create a new item inside storage.
func CreateItem(itemRequest models.ItemRequest) models.Item {
	// Create new Item to be passed to storage
	var newItem models.Item = models.Item{
		ID:        uuid.New().String(),
		Name:      itemRequest.Name,
		Price:     itemRequest.Price,
		Quantity:  itemRequest.Quantity,
		CreatedAt: time.Now(),
	}

	// Store the created item into storage
	storage = append(storage, newItem)

	// Return the new Item created
	return newItem
}

// UpdateItem : function to update the item’s data by ID.
func UpdateItem(itemRequest models.ItemRequest, id string) (models.Item, error) {
	// Iterate thru all the items int storage
	for index, item := range storage {
		// Check it the id parameter match the current item's ID
		if id == item.ID {
			// If it does, then update the item's data
			item.Name = itemRequest.Name
			item.Price = itemRequest.Price
			item.Quantity = itemRequest.Quantity
			item.UpdatedAt = time.Now()

			// Now replace the item in the storage with new update
			storage[index] = item

			// return the updated item and nil
			return item, nil
		}
	}
	// If no item was found, return empty item and error
	return models.Item{}, errors.New("item was not found")
}
