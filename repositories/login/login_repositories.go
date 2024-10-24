package login_repositories

import (
	"database/sql"
	"fmt"
	login_models "go-chat/models/login"
)

func GetLoginByUsername(db *sql.DB, username string) (*login_models.Login, error) {

	var login login_models.Login

	query := `
	SELECT *
	FROM logins
	WHERE username=$1
	`
	row := db.QueryRow(query, username)
	err := row.Scan(&login.Username, &login.Password, &login.IDUser)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("login not found for username: %s", username)
		}
		return nil, err
	}

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
