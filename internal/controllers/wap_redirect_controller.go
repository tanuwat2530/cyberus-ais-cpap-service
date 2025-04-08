package controllers

import (
	services "CyberusAisCpapBackend/internal/services"
	utils "CyberusAisCpapBackend/internal/utils/response"
	

	"net/http"
)

func WapRedirect(w http.ResponseWriter, r *http.Request) {
	// Check if the method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := services.WapRedirectProcess(r)

	utils.ResponseWithJSON(w, http.StatusOK, response)
}
