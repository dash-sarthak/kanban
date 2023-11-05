package v1

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/dash-sarthak/kanban/internal/database"
	"github.com/dash-sarthak/kanban/models"
	"github.com/dash-sarthak/kanban/util"
	"github.com/google/uuid"
	"net/http"
)

func (cfg *apiConfig) handleCreateProject(w http.ResponseWriter, r *http.Request) {
	type Parameters struct {
		Name        string
		Author      string
		Description string
	}
	decoder := json.NewDecoder(r.Body)
	params := Parameters{}
	decodeErr := decoder.Decode(&params)
	if decodeErr != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Could not decode parameters")
	}

	authorID, err := cfg.DB.GetAuthorID(r.Context(), params.Author)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("User '%v' does not exist", params.Author))
		return
	}

	var projectDescription sql.NullString
	if params.Description == "" {
		projectDescription = sql.NullString{String: params.Description, Valid: false}
	} else {
		projectDescription = sql.NullString{String: params.Description, Valid: true}
	}

	project, prErr := cfg.DB.CreateProject(r.Context(), database.CreateProjectParams{
		ID:          uuid.New(),
		Name:        params.Name,
		Author:      authorID,
		Description: projectDescription,
	})

	if prErr != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create project -> %v", prErr))
	}
	util.RespondWithJSON(w, http.StatusOK, models.MarshalProject(project))
}

//func (cfg *apiConfig) handleGetProject(w http.ResponseWriter, r *http.Request) {
//	author := r.UR
//	decodeErr := decoder.Decode(&params)
//	if decodeErr != nil {
//		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Could not decode parameters -> %v", decodeErr))
//	}
//
//	authorID, err := cfg.DB.GetAuthorID(r.Context(), params.Author)
//	if err != nil {
//		util.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("User '%v' does not exist", params.Author))
//		return
//	}
//
//	var projectDescription sql.NullString
//	if params.Description == "" {
//		projectDescription = sql.NullString{String: params.Description, Valid: false}
//	} else {
//		projectDescription = sql.NullString{String: params.Description, Valid: true}
//	}
//
//	project, prErr := cfg.DB.CreateProject(r.Context(), database.CreateProjectParams{
//		ID:          uuid.New(),
//		Name:        params.Name,
//		Author:      authorID,
//		Description: projectDescription,
//	})
//
//	if prErr != nil {
//		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create project -> %v", prErr))
//	}
//	util.RespondWithJSON(w, http.StatusOK, project)
//}
