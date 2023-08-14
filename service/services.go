package service

import (
	"errors"
	"github.com/BrightOlawale/Inventory-API/database"
	"github.com/BrightOlawale/Inventory-API/models"
	"github.com/google/uuid"
	"time"
)

// Create variable to store item data
// This is our storage, we will use this to store all the items
// It is stored in memory, so when the application is restarted, all the data will be lost
//var storage []models.Item = []models.Item{}

// GetAllItems : get all items from the storage.
func GetAllItems() []models.Item {
	// Create variable to store Item data
	var itemData []models.Item = []models.Item{}

	// Get all the data from the DB order by created_at
	database.DB.Order("created_at desc").Find(&itemData)

	// Return the itemData
	return itemData
}

// GetItemById : get the item’s data by ID.
func GetItemById(id string) (models.Item, error) {
	// Create a variable to store item data
	var itemData models.Item

	// Get item from the DB by id
	result := database.DB.First(&itemData, "id = ?", id)

	// If the item was not found, return an error
	if result.RowsAffected == 0 {
		return models.Item{}, errors.New("no item found")
	}

	// Return item since it was found
	return itemData, nil
}

// CreateItem : function to create a new item inside storage.
func CreateItem(itemRequest models.ItemRequest) models.Item {
	// Create new Item to be passed to the DB
	var newItem models.Item = models.Item{
		ID:        uuid.New().String(),
		Name:      itemRequest.Name,
		Price:     itemRequest.Price,
		Quantity:  itemRequest.Quantity,
		CreatedAt: time.Now(),
	}

	// Store the created item into DB
	database.DB.Create(&newItem)

	// Return the new Item created
	return newItem
}

// UpdateItem : function to update the item’s data by ID.
func UpdateItem(itemRequest models.ItemRequest, id string) (models.Item, error) {
	// Retrieve the item ID
	item, err := GetItemById(id)

	// If an error occured
	if err != nil {
		return models.Item{}, err
	}

	// Now update item data
	item.Name = itemRequest.Name
	item.Price = itemRequest.Price
	item.Quantity = itemRequest.Quantity
	item.UpdatedAt = time.Now()

	// Update the item in DB
	database.DB.Save(&item)

	// Return the updated item
	return item, nil
}

// DeleteItem : Function to delete items by ID.
func DeleteItem(id string) bool {
	// Get the item Data by ID
	itemData, err := GetItemById(id)

	if err != nil {
		return false
	}

	// Delete itemData from db
	database.DB.Delete(&itemData)

	// Return true to indicate deletion was successfully
	return true
}
