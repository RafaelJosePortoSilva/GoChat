package login_models

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IDUser   string `json:"IDUser"`
}
