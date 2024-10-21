package routes

import (
	login "go-chat/handlers/login"

	"github.com/gorilla/mux"
)

func SetupRouters() {
	router := mux.NewRouter()

	router.HandleFunc("/login/auth", login.HandleLogin)
}
