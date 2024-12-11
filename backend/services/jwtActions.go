package services

import (
	"Ariadne_Management/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// JWT Secret Key (you should keep this key secret and store it securely)
var jwtSecret = []byte("yourSecretKey")

// GenerateJWT generates a JWT token for a user
func GenerateJWT(user *models.User) (string, error) {
	// Define token expiration time
	expirationTime := time.Now().Add(1 * time.Hour)

	// Create the JWT claims, including user ID in the claims
	claims := &jwt.RegisteredClaims{
		Subject:   user.Username,
		ID:        strconv.Itoa(user.ID), // Store user ID in the ID field
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	// Create a new JWT token using the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using your secret key
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// ValidateJWT validates the JWT token and returns the claims if valid
func ValidateJWT(tokenString string) (*jwt.RegisteredClaims, error) {
	// Parse and validate the JWT token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	// Extract claims
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

// ExtractUserIDFromToken extracts the user ID from the JWT token
func ExtractUserIDFromToken(tokenString string) (int, error) {
	// Remove the "Bearer " prefix if present
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Validate the token and extract claims
	claims, err := ValidateJWT(tokenString)
	if err != nil {
		return 0, err
	}

	// Extract the user ID from the claims
	userID, err := strconv.Atoi(claims.ID)
	if err != nil {
		return 0, fmt.Errorf("invalid user ID in token")
	}

	return userID, nil
}

// AuthenticateJWT middleware to validate JWT token and authenticate the user
func AuthenticateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		// Validate and extract user ID from token
		userID, err := ExtractUserIDFromToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set the user ID in the context (can be accessed in further handlers)
		c.Set("userID", userID)

		c.Next()
	}
}
