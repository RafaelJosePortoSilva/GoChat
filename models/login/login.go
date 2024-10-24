package login_models

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IDUser   string `json:"IDUser"`
}

func CreateTableLogins() string {
	return `
CREATE TABLE IF NOT EXISTS logins (
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(50) NOT NULL,
    id_user INT REFERENCES users(id) NOT NULL
);
	`
}
