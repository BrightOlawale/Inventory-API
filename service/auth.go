package service

import (
	"errors"
	"github.com/BrightOlawale/Inventory-API/database"
	"github.com/BrightOlawale/Inventory-API/models"
	"github.com/BrightOlawale/Inventory-API/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Signup : Return JWT Token for user
func Signup(userDetails models.UserRequest) (string, error) {
	// Create a password using bcrypt library
	password, err := bcrypt.GenerateFromPassword([]byte(userDetails.Password), bcrypt.DefaultCost)

	// If error occurred
	if err != nil {
		return "", err
	}

	// Create a new user object that will be added to the db
	var newUser models.User = models.User{
		ID:       uuid.New().String(),
		Email:    userDetails.Email,
		Password: string(password),
	}

	// Create the user in the database
	database.DB.Create(&newUser)

	// Generate the JWT Token for user
	token, err := utils.GenerateNewAccessToken()

	// If error occurred
	if err != nil {
		return "", err
	}

	// Return the JTW token
	return token, nil
}

// Login : Return JWT Token for user
func Login(userDetails models.UserRequest) (string, error) {
	// Create a variable to hold the user data
	var user models.User

	// Get the user from the database using the email
	queryResult := database.DB.Where("email = ?", userDetails.Email).First(&user)

	// If user was not found
	if queryResult.Error != nil {
		return "", errors.New("user not found")
	}

	// Compare the password from the request with the password from the database
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDetails.Password))

	// If password does not match
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate the JWT Token for user
	token, err := utils.GenerateNewAccessToken()

	// If token generation failed
	if err != nil {
		return "", err
	}

	// Return the JTW token
	return token, nil
}
