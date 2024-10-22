package services_login

import (
	"fmt"
	login_repo "go-chat/repositories/login"

	"golang.org/x/crypto/bcrypt"
)

func AuthUser(username string, pass string) error {

	err := login_repo.VerifyUsernameExists(username)
	if err != nil {
		return fmt.Errorf("invalid username")
	}

	hashPassInDB, err := login_repo.GetPasswordByUsername(username)
	if err != nil {
		return fmt.Errorf("internal error")
	}

	isHashCorrect := checkPasswordHash(pass, hashPassInDB)

	if isHashCorrect {
		return nil
	} else {
		return fmt.Errorf("invalid password")
	}

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
