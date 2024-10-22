package user_services

import (
	"fmt"
	chat_models "go-chat/models/chat"
	chat_repo "go-chat/repositories/chat"
)

func GetUser(id string) (*chat_models.User, error) {

	user, err := chat_repo.GetUserById(id)
	if err != nil {
		return nil, fmt.Errorf("invalid username")
	}
	return user, nil

}