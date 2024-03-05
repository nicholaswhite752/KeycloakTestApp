package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Nerzal/gocloak/v7"
)

type keycloak struct {
	gocloak      gocloak.GoCloak // keycloak client
	clientId     string          // clientId specified in Keycloak
	clientSecret string          // client secret specified in Keycloak
	realm        string          // realm specified in Keycloak
}

func newKeycloak() *keycloak {
	return &keycloak{
		gocloak:      gocloak.NewClient("http://localhost:8080"),
		clientId:     "test-react",
		clientSecret: "elG1v2vJ4Ql9NY9I1iMzMZMGDio3GFFS",
		realm:        "TestRealm",
	}
}

func (kcInstance *keycloak) extractBearerToken(token string) string {
	return strings.Replace(token, "Bearer ", "", 1)
}

func (kcInstance *keycloak) verifyToken(next http.Handler) http.Handler {

	f := func(w http.ResponseWriter, r *http.Request) {
		// try to extract token from cookie, it is called token on frontend
		cookie, _ := r.Cookie("token")
		if cookie.Valid() != nil {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		token := cookie.Value

		if token == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		if token == "" {
			http.Error(w, "Bearer Token missing", http.StatusUnauthorized)
			return
		}

		//// call Keycloak API to verify the access token
		result, err := kcInstance.gocloak.RetrospectToken(context.Background(), token, kcInstance.clientId, kcInstance.clientSecret, kcInstance.realm)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid or malformed token: %s", err.Error()), http.StatusUnauthorized)
			return
		}

		_, _, err = kcInstance.gocloak.DecodeAccessToken(context.Background(), token, kcInstance.realm, "")
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid or malformed token: %s", err.Error()), http.StatusUnauthorized)
			return
		}

		//jwtj, _ := json.Marshal(jwt)
		//fmt.Printf("token: %v\n", string(jwtj))

		// check if the token isn't expired and valid
		if !*result.Active {
			http.Error(w, "Invalid or expired Token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(f)
}