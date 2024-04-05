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
	mux.Post("/greet/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		w.Write([]byte(fmt.Sprintf("hello %s\n", name)))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")), mux))
}
