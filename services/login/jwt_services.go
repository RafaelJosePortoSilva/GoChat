package services_login

import (
	"fmt"
	chat_models "go-chat/models/chat"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("your-secret-key")

func GenerateJWT(user *chat_models.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":             user.ID,                          // Subject (user identifier)
		"name":           user.Name,                        // Issuer
		"expirationTime": time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat":            time.Now().Unix(),                // Issued at
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWTToken(tokenString string) (*jwt.Token, error) {
	// Parse the token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// Check for verification errors
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Return the verified token
	return token, nil
}

func GetUserIDFromJWTToken(token *jwt.Token) (string, error) {

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["ID"] == nil {
		return "", fmt.Errorf("invalid token")
	}
	return claims["ID"].(string), nil
}
