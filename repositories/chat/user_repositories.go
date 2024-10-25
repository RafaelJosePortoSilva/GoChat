package chat_repositories

import (
	"database/sql"
	chat_models "go-chat/models/chat"
)

func GetUserById(db *sql.DB, id string) (*chat_models.User, error) {

	return nil, nil

}

func CreateNewUser(db *sql.DB, id string) error {

	var aux string

	query := `INSERT INTO logins (id_login) VALUES ($1) RETURNING id_login`
	err := db.QueryRow(query, id).Scan(&aux)
	if err != nil {
		return err
	}

	return nil

}
