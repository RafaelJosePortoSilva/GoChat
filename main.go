package main

import (
	"fmt"
	"go-chat/database"
	"go-chat/routes"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	// Buscando variaveis de ambiente
	// dbName, dbUser, dbPassword, dbHost, dbPort

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading env file: %v", err)
	}

	// Acessa as variáveis de ambiente
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// Configurando base de dados
	db, err := database.OpenConn(dbName, dbUser, dbPassword, dbHost, dbPort)
	if err != nil {
		log.Fatalf("Database open failed: %v\n", err)
	}

	err = database.CreateDatabaseIfNotExists(dbName, dbUser, dbPassword, dbHost, dbPort)
	if err != nil {
		log.Fatalf("Database creating failed: %v\n", err)
	}

	err = database.CreateTablesIfNotExists(db)
	if err != nil {
		log.Fatalf("Database tables creating failed: %v\n", err)
	}

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

	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v\n", err)
	}
	fmt.Printf("Address: %s", addr)
}
