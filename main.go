package main

import (
	"fmt"
	"go-chat/database"
	"go-chat/routes"
	"log"
	"net/http"
	"time"
)

func main() {

	// Buscando variaveis de ambiente

	// Configurando base de dados
	db := database.OpenConn()

	// Configurando servidor
	r := routes.SetupRouters()
	addr := ":8080"
	fmt.Printf("Setup routes OK\n")

	srv := &http.Server{
		Handler:           r,
		Addr:              addr,
		WriteTimeout:      15 * time.Second,
		ReadTimeout:       15 * time.Second,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}
	fmt.Printf("Server created\n")

	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v\n", err)
	}
	fmt.Printf("Address: %s", addr)
}
