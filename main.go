package main

import (
	"go-chat/routes"
	"log"
	"net/http"
	"time"
)

func main() {
	r := routes.SetupRouters()
	addr := ":8080"

	srv := &http.Server{
		Handler:           r,
		Addr:              addr,
		WriteTimeout:      15 * time.Second,
		ReadTimeout:       15 * time.Second,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
