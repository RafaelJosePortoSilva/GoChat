package routes

import (
	login_handlers "go-chat/handlers/login"

	"github.com/gorilla/mux"
)

func SetupRouters() *mux.Router {
	r := mux.NewRouter()

	chat := r.PathPrefix("/chat").Subrouter()
	chat.Use(login_handlers.AuthMiddleware)

	r.HandleFunc("/login/auth", login_handlers.HandleLogin)

	return r
}
