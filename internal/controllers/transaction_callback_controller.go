package controllers

import (
	services "cyberus/ais-cacp-service/internal/services"
	utils "cyberus/ais-cacp-service/internal/utils/response"

	"net/http"
)

func TransactionCallback(w http.ResponseWriter, r *http.Request) {
	// Check if the method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := services.TransactionCallbackProcessRequest(r)

	utils.ResponseWithJSON(w, http.StatusOK, response)
}
