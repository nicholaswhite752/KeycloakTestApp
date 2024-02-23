package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
)

//	@title			Arcus API
//	@version 		1.0
//	@description 	API backend for the Arcus project

// 	@host 			localhost:5050
// 	@BasePath 		/api/v1
//	@accept 		json

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Keycloak Access Key Authorization


func main() {
	router := createRouter()

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	http.ListenAndServe(":3000", router)
}
