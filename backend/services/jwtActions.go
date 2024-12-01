package services

import (
	"Ariadne_Management/models"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

// GenerateJWT generates a JWT token for a user
func GenerateJWT(user *models.User) (string, error) {
	// Define token expiration time
	expirationTime := time.Now().Add(1 * time.Hour)

	// Create the JWT claims, including user ID in the claims
	claims := &jwt.RegisteredClaims{
		Subject:   user.Username,
		ID:        strconv.Itoa(user.ID),
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	// Create a new JWT token using the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using your secret key
	signedToken, err := token.SignedString([]byte("yourSecretKey"))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// ExtractUserIDFromToken parces the token to get the ID from it
func ExtractUserIDFromToken(tokenString string) (int, error) {
	// Parse and validate the JWT token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("yourSecretKey"), nil
	})
	if err != nil {
		return 0, err
	}

	// Extract claims
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	// Extract the user ID from the 'jti' field
	userID, err := strconv.Atoi(claims.ID)
	if err != nil {
		return 0, fmt.Errorf("invalid user ID in token")
	}

	return userID, nil
}
