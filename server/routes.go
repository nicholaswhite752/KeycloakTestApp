package main

import "github.com/go-chi/chi"

func createRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", MainHandler)
	router.Get("/refreshToken", RefreshTokenHandler)

	return router
}