package services_login

import (
	"fmt"
	chat_models "go-chat/models/chat"
	login_repo "go-chat/repositories/login"
	user_services "go-chat/services/chat"

	"golang.org/x/crypto/bcrypt"
)

func AuthUser(username string, pass string) (*chat_models.User, error) {

	login, err := login_repo.GetLoginByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("invalid username")
	}

	hashPassInDB := login.Password
	if err != nil {
		return nil, fmt.Errorf("internal error")
	}

	isHashCorrect := checkPasswordHash(pass, hashPassInDB)

	if !isHashCorrect {
		return nil, fmt.Errorf("invalid password")
	}

	user, err := user_services.GetUser(login.IDUser)
	if err != nil {
		return nil, fmt.Errorf("invalid user")
	}

	return user, nil

}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
