package routes

import (
	"CyberusAisCpapBackend/internal/controllers"
	"net/http"
)

// SetupRoutes registers all application routes
func SetupRoutes() {
	// Register routes using http.HandleFunc
	http.HandleFunc("/api/wap-redirect", controllers.WapRedirect)
	http.HandleFunc("/api/subscription-callback", controllers.SubscriptionCallback)
	http.HandleFunc("/api/transaction-callback", controllers.TransactionCallback)
	http.HandleFunc("/api/", HomeHandler)
}

// HomeHandler for root endpoint
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to backend API power by GoLang ^_^"))
}
