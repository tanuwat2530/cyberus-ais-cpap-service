package main

import (
	"CyberusAisCpapBackend/internal/routes"
	"fmt"
	"net/http"
)

func main() {
	// Initialize database

	// config.InitDB()
	// defer config.DB.Close()

	// Setup all routes
	routes.SetupRoutes()

	// Start the server on port 8080
	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
