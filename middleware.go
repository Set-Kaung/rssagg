package main

import (
	"log"
	"net/http"

	"github.com/Set-Kaung/rssagg/internal/auth"
	"github.com/Set-Kaung/rssagg/internal/database"
)

type authentication func(http.ResponseWriter, *http.Request, database.User)

func (ap *apiConfig) authenticationMiddleware(authed authentication) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, "auth error "+err.Error())
			return
		}
		user, err := ap.DB.GetUserByAPIKey(r.Context(), apikey)
		if err != nil {
			log.Println(err)
			respondWithError(w, 500, "couldn't get user")
			return
		}
		authed(w, r, user)
	}
}
