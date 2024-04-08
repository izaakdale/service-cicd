package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/izaakdale/ittp"
)

func main() {
	mux := ittp.NewServeMux()
	mux.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	mux.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("healthy"))
	})

	mux.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("updated test"))
	})

	mux.Post("/greet/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		w.Write([]byte(fmt.Sprintf("hello %s\n", name)))
	})

	mux.Get("/new", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("new version handler"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	log.Printf("up and running\n")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("HOST"), port), mux))
}
