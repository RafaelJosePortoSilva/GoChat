package login_handlers

import (
	"context"
	"fmt"
	login_services "go-chat/services/login"
	"net/http"
	"strings"
)

type contextKey string

const UserIDContextKey = contextKey("id")

// Middleware para verificar o JWT
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("AuthMiddleware: entrou no middleware")

		// Obtem o token do cabeçalho
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Verifica se o token não está vazio
		if tokenString == "" {
			http.Redirect(w, r, "/login/", http.StatusSeeOther)
			fmt.Printf("AuthMiddleware: tokenString vazia: %s", tokenString)
			return
		}

		// Verify the token
		token, err := login_services.VerifyJWTToken(tokenString)
		if err != nil {
			fmt.Printf("Token verification failed: %v\\n", err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		userId, err := login_services.GetUserIDFromJWTToken(token)
		if err != nil {
			fmt.Printf("Token verification failed: %v\\n", err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		ctx := context.WithValue(r.Context(), UserIDContextKey, userId)

		// Chama o próximo handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
