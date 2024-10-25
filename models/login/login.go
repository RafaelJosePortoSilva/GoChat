package login_models

type Login struct {
	ID       string `json:"ID"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateTableLogins() string {
	return `
CREATE TABLE IF NOT EXISTS logins (
	id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(50) NOT NULL
);
	`
}
