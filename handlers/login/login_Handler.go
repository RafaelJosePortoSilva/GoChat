package login_handlers

import (
	"encoding/json"
	"fmt"
	login_models "go-chat/models/login"
	login_services "go-chat/services/login"
	"net/http"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {

	var login login_models.Login
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Bad Request"}`))
		return
	}

	err = login_services.AuthUser(login.Username, login.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprintf(`{"message": %s}`, err.Error())))
		return
	}

}
