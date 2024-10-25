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
	err := row.Scan(&login.Username, &login.Password, &login.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("login not found for username: %s", username)
		}
		return nil, err
	}

	return nil, nil
}

func CreateNewLogin(db *sql.DB, username string, hash string) (string, error) {

	var id string
	if hasUsernameDuplicity(db, username) {
		return "", fmt.Errorf("duplicated username")
	}

	query := `INSERT INTO logins (username, password) VALUES ($1, $2) RETURNING id`
	err := db.QueryRow(query, username, hash).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil

}

func hasUsernameDuplicity(db *sql.DB, username string) bool {

	var aux string
	query := `
	SELECT *
	FROM logins
	WHERE username=$1
	`
	err := db.QueryRow(query, username).Scan(&aux)
	return !(err == nil)
}

func DeleteLogin(db *sql.DB, id string) (int64, error) {

	query := `DELETE FROM logins WHERE id=$1`
	res, err := db.Exec(query, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()

}
