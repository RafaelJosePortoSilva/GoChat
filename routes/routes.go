package routes

import (
	login_handlers "go-chat/handlers/login"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouters() *mux.Router {
	r := mux.NewRouter()

	chat := r.PathPrefix("/chat").Subrouter()
	chat.Use(login_handlers.AuthMiddleware)

	login := r.PathPrefix("/login").Subrouter()

	login.HandleFunc("/auth", login_handlers.HandleLogin)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/login.html")
	})

	return r
}
