package chat_repositories

import (
	"database/sql"
	"fmt"
	chat_models "go-chat/models/chat"
)

func GetUserById(db *sql.DB, id string) (*chat_models.User, error) {

	var user chat_models.User

	query := `
	SELECT *
	FROM users
	WHERE id_login=$1
	`
	row := db.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Name, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("login not found for id: %s", id)
		}
		return nil, err
	}
	return &user, nil

}

func CreateNewUser(db *sql.DB, id, username string) error {

	var aux string

	query := `INSERT INTO users (id_login, Name) VALUES ($1, $2) RETURNING id_login`
	err := db.QueryRow(query, id, username).Scan(&aux)
	if err != nil {
		return err
	}

	return nil

}
