package login_repositories

import (
	"fmt"
	login_models "go-chat/models/login"
)

func GetLoginByUsername(username string) (*login_models.Login, error) {
	return nil, nil
}

func CreateNewLogin(username string, hash string) error {
	if !verifyUsernameDuplicity(username) {
		return nil
	} else {
		return fmt.Errorf("cannot create account")
	}
}

func verifyUsernameDuplicity(username string) bool {
	return false
}
