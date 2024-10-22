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

	user, err := login_services.AuthUser(login.Username, login.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprintf(`{"message": %s}`, err.Error())))
		return
	}

	// Colocar JWT AQUI

	// Armazena o ID do usuário na sessão
	http.SetCookie(w, &http.Cookie{
		Name:  "user_id",
		Value: user.ID, // Armazena o ID do usuário
		Path:  "/",
		// Configure outras opções conforme necessário, como HttpOnly, Secure, etc.
	})

	// Redireciona para a página do sistema
	http.Redirect(w, r, "/index", http.StatusSeeOther)

}
