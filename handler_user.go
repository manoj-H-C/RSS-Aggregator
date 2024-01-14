package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/manoj-H-C/rssagg/internal/auth"
	"github.com/manoj-H-C/rssagg/internal/database"
)

func (apicfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}
	user, err := apicfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}
	respondWithJSON(w, 200, databaseUserToUser(user))
}

func (apicfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apikey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Auth error: %s", err))
		return
	}
	user, err := apicfg.DB.GetUserByAPIKey(r.Context(), apikey)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't get user: %s", err))
		return
	}
	respondWithJSON(w, 200, databaseUserToUser(user))

}
