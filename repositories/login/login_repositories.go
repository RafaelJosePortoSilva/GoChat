package login_repositories

import (
	"database/sql"
	"fmt"
	login_models "go-chat/models/login"
)

func GetLoginByUsername(db *sql.DB, username string) (*login_models.Login, error) {
	return nil, nil
}

func CreateNewLogin(db *sql.DB, username string, hash string) error {
	if !verifyUsernameDuplicity(db, username) {
		return nil
	} else {
		return fmt.Errorf("cannot create account")
	}
}

func verifyUsernameDuplicity(db *sql.DB, username string) bool {
	return false
}
