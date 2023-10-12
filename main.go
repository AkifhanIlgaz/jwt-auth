package main

import (
	"net/http"

	"github.com/AkifhanIlgaz/jwt-auth/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Post("/api/auth/login", handlers.Login)

	http.ListenAndServe(":3000", r)
}
