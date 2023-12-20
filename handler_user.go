package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/KieranWebDev/Go-BlogAggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, req *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(req.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "error parsing json")
	}

	user, err := apiCfg.DB.CreateUser(req.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, "error creating user")
	}
	configuredUser := databaseUserToUser(user)
	respondWithJSON(w, http.StatusOK, configuredUser)
}
