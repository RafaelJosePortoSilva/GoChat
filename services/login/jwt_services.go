package services_login

import (
	chat_models "go-chat/models/chat"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("your-secret-key")

func GenerateJWT(user *chat_models.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":             user.ID,                          // Subject (user identifier)
		"username":       user.Username,                    // Issuer
		"expirationTime": time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat":            time.Now().Unix(),                // Issued at
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
