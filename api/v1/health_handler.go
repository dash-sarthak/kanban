package v1

import (
	"github.com/dash-sarthak/kanban/util"
	"net/http"
)

func (cfg *apiConfig) handleCheckHealth(w http.ResponseWriter, _ *http.Request) {
	type health struct {
		Health string
	}
	util.RespondWithJSON(w, http.StatusOK, health{Health: "Up"})
}
