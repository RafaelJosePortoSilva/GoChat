package routes

import (
	"fmt"
	login_handlers "go-chat/handlers/login"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouters() *mux.Router {
	r := mux.NewRouter()

	// Rota básica de teste
	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Test route is working")
	})

	// Serve arquivos estáticos da pasta "static/login/"
	r.PathPrefix("/static/login/").Handler(http.StripPrefix("/static/login/", http.FileServer(http.Dir("./static/login/"))))

	login := r.PathPrefix("/login").Subrouter()

	login.HandleFunc("/auth", login_handlers.HandleLogin)
	login.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/login/login.html")
	})

	chat := r.PathPrefix("/chat").Subrouter()
	chat.Use(login_handlers.AuthMiddleware)

	return r
}
