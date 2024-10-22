package login_handlers

import (
	"fmt"
	login_services "go-chat/services/login"
	"net/http"
	"strings"
)

// Middleware para verificar o JWT
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Obtem o token do cabeçalho
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Verifica se o token não está vazio
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Verify the token
		_, err := login_services.VerifyJWTToken(tokenString)
		if err != nil {
			fmt.Printf("Token verification failed: %v\\n", err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Chama o próximo handler
		next.ServeHTTP(w, r)
	}
}
