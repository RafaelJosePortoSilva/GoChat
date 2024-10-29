package services_login

import (
	"database/sql"
	"errors"
	"fmt"
	chat_models "go-chat/models/chat"
	login_repo "go-chat/repositories/login"
	user_services "go-chat/services/chat"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func AuthUser(db *sql.DB, username string, pass string) (*chat_models.User, error) {

	login, err := login_repo.GetLoginByUsername(db, username)
	if err != nil {
		return nil, err
	}

	if login == nil {
		return nil, fmt.Errorf("cannot find user")
	}

	hashPassInDB := login.Password

	isHashCorrect := checkPasswordHash(pass, hashPassInDB)

	if !isHashCorrect {
		return nil, fmt.Errorf("invalid password")
	}

	user, err := user_services.GetUser(db, login.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid user")
	}

	return user, nil

}

func CreateLogin(db *sql.DB, username string, pass string) error {

	err := verifyUsername(username)
	if err != nil {
		return err
	}

	err = verifyPassword(pass)
	if err != nil {
		return err
	}
	hash, err := hashPassword(pass)
	if err != nil {
		return err
	}

	id, err := login_repo.CreateNewLogin(db, username, hash)
	if err != nil {
		return err
	}

	// Agora, criar o user correspondente
	// Se não conseguir criar o user, apagar o login e cancelar a operação

	err = user_services.CreateUser(db, id, username)
	if err != nil {
		_, err := DeleteLogin(db, id)
		if err != nil {
			return err
		}
		return err
	}

	return nil

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

func verifyUsername(username string) error {
	var validUsername = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if !validUsername.MatchString(username) {
		return fmt.Errorf("invalid characters")
	}
	if len(username) > 50 {
		return fmt.Errorf("so long username")
	}
	return nil
}

func verifyPassword(password string) error {
	// Verifica o tamanho mínimo
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	// Expressões regulares para cada critério
	hasUpperCase := regexp.MustCompile(`[A-Z]`).MatchString
	hasLowerCase := regexp.MustCompile(`[a-z]`).MatchString
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString
	hasSpecialChar := regexp.MustCompile(`[!@#\$%\^&\*]`).MatchString

	// Verifica se a senha contém uma letra maiúscula
	if !hasUpperCase(password) {
		return errors.New("password must contain at least one uppercase letter")
	}

	// Verifica se a senha contém uma letra minúscula
	if !hasLowerCase(password) {
		return errors.New("password must contain at least one lowercase letter")
	}

	// Verifica se a senha contém um número
	if !hasNumber(password) {
		return errors.New("password must contain at least one number")
	}

	// Verifica se a senha contém um caractere especial
	if !hasSpecialChar(password) {
		return errors.New("password must contain at least one special character (!@#$%^&*)")
	}

	// Se todas as verificações passaram
	return nil
}

func DeleteLogin(db *sql.DB, id string) (int64, error) {

	rows_affected, err := login_repo.DeleteLogin(db, id)
	if err != nil {
		return 0, err
	}
	return rows_affected, nil

}
