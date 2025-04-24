package routes

import (
	"cyberus/ais-cacp-service/internal/controllers"
	"net/http"
)

// SetupRoutes registers all application routes
func SetupRoutes() {
	// Register routes using http.HandleFunc
	http.HandleFunc("/ais/wap-redirect", controllers.WapRedirect)
	http.HandleFunc("/ais/subscription-callback", controllers.SubscriptionCallback)
	http.HandleFunc("/ais/transaction-callback", controllers.TransactionCallback)
	http.HandleFunc("/ais/", HomeHandler)
}

// HomeHandler for root endpoint
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to backend API power by GoLang ^_^"))
}
