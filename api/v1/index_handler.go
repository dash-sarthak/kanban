package v1

import (
	"net/http"
)

func (cfg *apiConfig) handleIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/v1/home", http.StatusMovedPermanently)
	return
}
