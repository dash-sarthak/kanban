package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dash-sarthak/kanban/internal/database"
	"github.com/dash-sarthak/kanban/util"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handleAuthorsCreate(w http.ResponseWriter, r *http.Request) {
	type Parameters struct {
		Name     string
		Username string
	}
	decoder := json.NewDecoder(r.Body)
	params := Parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Could not decode parameters")
	}
	author, err := cfg.DB.CreateAuthor(r.Context(), database.CreateAuthorParams{
		ID:       uuid.New(),
		Name:     params.Name,
		Username: params.Username,
	})
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Could not create author -> %v", err))
	}
	util.RespondWithJSON(w, http.StatusOK, author)
}

func (cfg *apiConfig) handleAuthorsFetch(w http.ResponseWriter, r *http.Request) {
	authors, err := cfg.DB.GetAuthors(r.Context())
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Could not fetch authors -> %v", err))
	}
	util.RespondWithJSON(w, http.StatusOK, authors)
}

func (cfg *apiConfig) handleAuthorsDelete(w http.ResponseWriter, r *http.Request) {

	authors, err := cfg.DB.GetAuthors(r.Context())
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Could not fetch authors -> %v", err))
	}
	util.RespondWithJSON(w, http.StatusOK, authors)
}
