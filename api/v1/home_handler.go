package v1

import (
	"fmt"
	"github.com/dash-sarthak/kanban/util"
	"html/template"
	"net/http"
)

func (cfg *apiConfig) handleHome(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("template/index.html"))
	err := templ.Execute(w, nil)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Could not render template -> %v", err))
	}
}
