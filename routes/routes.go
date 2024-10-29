package routes

import (
	"database/sql"
	"fmt"
	chat_handlers "go-chat/handlers/chat"
	login_handlers "go-chat/handlers/login"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouters(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	// Rota básica de teste
	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Test route is working")
	})

	// Serve arquivos estáticos da pasta "static/login/"
	r.PathPrefix("/static/login/").Handler(http.StripPrefix("/static/login/", http.FileServer(http.Dir("./static/login/"))))
	r.PathPrefix("/static/newLogin/").Handler(http.StripPrefix("/static/newLogin/", http.FileServer(http.Dir("./static/newLogin/"))))

	login := r.PathPrefix("/login").Subrouter()

	login.HandleFunc("/auth", login_handlers.HandleLogin(db)).Methods("POST")
	login.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/login/login.html")
	}).Methods("GET")
	login.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login/", http.StatusSeeOther)
	})

	login.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/newLogin/newLogin.html")
	}).Methods("GET")
	login.HandleFunc("/new", login_handlers.HandleCreateLogin(db)).Methods("POST")

	chat := r.PathPrefix("/chat").Subrouter()
	chat.Use(login_handlers.AuthMiddleware)

	chat.HandleFunc("/", chat_handlers.HandleTest(db)).Methods("GET")

	return r
}
