package main

import (
	"net/http"
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
	router := CreateRouter()

	http.ListenAndServe(":5050", router)
}
