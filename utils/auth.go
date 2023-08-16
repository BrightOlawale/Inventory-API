package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"strings"
	"time"
)

// JWTTokenMetaData JWT Token metadata
type JWTTokenMetaData struct {
	// The Expiration time of the token
	ExpiresAt int64 `json:"expires_at"`
}

// GenerateNewAccessToken : Generate and returns JWT Token
func GenerateNewAccessToken() (string, error) {
	// Get JTW Secret key from .env file
	secret := GetValue("JWT_SECRET")

	// Get the JTW Expiry time from .env file
	// strconv.Atoi() converts a string to an integer
	expiryPeriod, _ := strconv.Atoi(GetValue("JWT_EXPIRES_IN"))

	// Create JTW claim object
	// Claims are the information we want to store in the token and we can access them from the token
	claims := jwt.MapClaims{}

	// Add expiration time for the token
	// The Add method called on the time.Now() object adds the specified duration to the current time
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(expiryPeriod)).Unix()

	// Create a JWT Token with the JWT Claim object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Convert created token into a string format
	strToken, err := token.SignedString([]byte(secret))

	// Check for error
	if err != nil {
		return "", err
	}

	// Return the token and no error
	return strToken, nil
}

// ExtractTokenMetadata : Extracts the JWT Token metadata
func ExtractTokenMetadata(c *fiber.Ctx) (*JWTTokenMetaData, error) {
	// First, verify the token
	token, err := VerifyToken(c)

	// If verification failed, return error
	if err != nil {
		return nil, err
	}

	// Get the claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)

	// If the claims are not valid, return error
	if !ok || !token.Valid {
		return nil, err
	}

	// Set the expiry time
	expiryTime := int64(claims["exp"].(float64))

	// Return the token metadata
	return &JWTTokenMetaData{
		ExpiresAt: expiryTime,
	}, nil
}

// CheckToken : Check if the JWT Token is valid
func CheckToken(c *fiber.Ctx) (bool, error) {
	// Get current time
	currentTime := time.Now().Unix()

	// Extract the token claim data
	tokenClaims, err := ExtractTokenMetadata(c)

	// If we cannot get a valid token claim data, return false
	if err != nil {
		return false, err
	}

	// Get the expiry time from the token claim data
	expiryTime := tokenClaims.ExpiresAt

	// Check if the token has expired by comparing the current time and the expiry time
	if expiryTime < currentTime {
		return false, nil
	}

	// Return true if token is valid
	return true, nil
}

// ExtractToken : Extracts the JWT Token from the request header
func ExtractToken(c *fiber.Ctx) string {
	// Get the Authorization header from the request
	bearerToken := c.Get("Authorization")

	// Split the token from the bearer string
	// The token is the second element in the array
	token := strings.Split(bearerToken, " ")

	if len(token) == 2 {
		return token[1]
	}

	// Return an empty string if the token is not found
	return ""
}

// VerifyToken : Verify the JWT Token
func VerifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	// Get the token from the bearer token
	tokenString := ExtractToken(c)

	// Verify the token using the JWT Secret key
	// The jwt.Parse() method returns a token object and an error
	token, err := jwt.Parse(tokenString, jwtKeyFunc)

	// If verification failed, return error
	if err != nil {
		return nil, err
	}

	// Return the token and no error
	return token, nil
}

// jwtKeyFunc : JWT Key function, returns the JWT Secret key
func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(GetValue("JWT_SECRET")), nil
}

// NB: The interface{} type is the empty interface type.
// It is used to represent any type in Go.
// []byte is a byte slice. When used to wrap a string, it converts the string to a byte slice.
