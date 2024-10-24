package chat_models

import "time"

type User struct {
	ID        string    `json:"ID"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateTableUsers() string {
	return `
CREATE TABLE IF NOT EXISTS "users" (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`
}
