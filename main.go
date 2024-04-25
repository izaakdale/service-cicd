package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/izaakdale/ittp"
	"github.com/rs/cors"
)

func main() {
	mux := ittp.NewServeMux()
	mux.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	mux.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("healthy"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	mux.AddMiddleware(cors.Handler)

	log.Printf("up and running\n")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("HOST"), port), mux))
}
