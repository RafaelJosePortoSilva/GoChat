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
		fmt.Printf("Bad Request\n")
		return
	}

	user, err := login_services.AuthUser(login.Username, login.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprintf(`{"message": %s}`, err.Error())))
		fmt.Printf("Error login handler - auth user: %s\n", err.Error())
		return
	}

	// Gera o token JWT
	token, err := login_services.GenerateJWT(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Error generating token"}`))
		fmt.Printf("Error login handler - gen jwt: %s\n", err.Error())
		return
	}

	// Define o token no cabeçalho da resposta
	w.Header().Set("Authorization", "Bearer "+token)

	// Opcional: Retorna uma resposta JSON para confirmar que o login foi bem-sucedido
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"message": "Login successful", "token": "%s"}`, token)))
	fmt.Printf("Successful login")

	// Descomentar depois de criar o index
	//http.Redirect(w, r, "/index", http.StatusSeeOther)

}
