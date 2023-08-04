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
