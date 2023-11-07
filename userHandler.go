package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Set-Kaung/rssagg/internal/auth"
	"github.com/Set-Kaung/rssagg/internal/database"
	"github.com/google/uuid"
)

func (ap *apiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Println(err)
		respondWithError(w, 400, "Error parsing json")
		return
	}
	user, err := ap.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		log.Println(err)
		respondWithError(w, 500, "couldn't create user")
		return
	}
	respondWithJSON(w, 201, dbUsertoUser(user))
}

func (ap *apiConfig) GetUser(w http.ResponseWriter, r *http.Request) {
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
	respondWithJSON(w, 200, dbUsertoUser(user))
}
