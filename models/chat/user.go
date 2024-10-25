package chat_models

import "time"

type User struct {
	ID        string    `json:"IDLogin"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateTableUsers() string {
	return `
CREATE TABLE IF NOT EXISTS users (
	id_login INT REFERENCES logins(id) NOT NULL,
    name VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`
}
