package main

import (
	"REST-API/menu"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	menu.Register(router, nil)

	const port = ":8000"

	log.Println("server running at port", port)

	http.ListenAndServe(port, router)
}
