package user_services

import (
	"database/sql"
	"fmt"
	chat_models "go-chat/models/chat"
	chat_repo "go-chat/repositories/chat"
)

func GetUser(db *sql.DB, id string) (*chat_models.User, error) {

	user, err := chat_repo.GetUserById(db, id)
	if err != nil {
		return nil, fmt.Errorf("invalid username")
	}
	return user, nil

}

func CreateUser(db *sql.DB, id string, username string) error {

	err := chat_repo.CreateNewUser(db, id, username)
	if err != nil {
		return err
	}
	return nil

}
