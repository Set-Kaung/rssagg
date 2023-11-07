package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Set-Kaung/rssagg/internal/database"
	"github.com/google/uuid"
)

func (ap *apiConfig) CreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Println(err)
		respondWithError(w, 400, "Error parsing json")
		return
	}
	feed, err := ap.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		log.Println(err)
		respondWithError(w, 500, "couldn't create feed")
		return
	}
	respondWithJSON(w, 201, dbFeedtoFeed(feed))
}

func (ap *apiConfig) GetFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, dbUsertoUser(user))
}

func (ap *apiConfig) GetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := ap.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 500, "cannot get feeds")
		return
	}
	respondWithJSON(w, 200, dbFeedsToFedds(feeds))
}
