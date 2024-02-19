package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/millz147/rssagg/internal/auth"
	"github.com/millz147/rssagg/internal/database"
)

func (dbQuery *dbRepoStructure) handleCreateUser(w http.ResponseWriter, req *http.Request) {

	type parameter struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(req.Body)
	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON, %v", err))
		return
	}

	user, err := dbQuery.DB.CreateUser(req.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error Creating User, %v", err))
		return
	}

	responseWithJson(w, 200, convertDbUserToUser(user))
}

func (dbQuery *dbRepoStructure) handleGetUser(w http.ResponseWriter, req *http.Request) {

	api_key, err := auth.GetApikey(req.Header)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Auth Error %v ", err))
		return
	}

	user, err := dbQuery.DB.GetUserByApikey(req.Context(), api_key)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("invalid apikey %v ", err))
		return
	}

	responseWithJson(w, 202, convertDbUserToUser(user))

}
