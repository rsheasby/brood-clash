package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	fmt.Println("Starting server.")
	r := chi.NewRouter()
	r.Post("/login", Login)
	http.ListenAndServe(":3000", r)
}
