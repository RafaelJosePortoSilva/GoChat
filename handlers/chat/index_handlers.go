package chat_handlers

import (
	"database/sql"
	"fmt"
	login_handlers "go-chat/handlers/login"
	"net/http"
)

func HandleTest(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, ok := r.Context().Value(login_handlers.UserIDContextKey).(string)
		if !ok {
			http.Error(w, "User not found in context", http.StatusUnauthorized)
			return
		}

		fmt.Fprintf(w, "ID: %s!", id)

	}
}
