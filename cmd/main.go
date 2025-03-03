// cmd/main.go
package main

import (
	"log"
	"net/http"
	"pokemon-evolution/config"
	"pokemon-evolution/middleware"
	"pokemon-evolution/routes"

	_ "pokemon-evolution/docs"

	"github.com/gorilla/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Pokemon Evolution API
// @version 1.0
// @description Pokemon Evolution API .

// @host localhost:5001
// @BasePath /api

func main() {
	cfg := config.LoadConfig()
	router := routes.RegisterRoutes()

	// Apply logging middleware
	loggedRouter := middleware.LoggingMiddleware(router)

	// Enable CORS
	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)
	//http://localhost:5000/swagger/index.html
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Printf("Server running on IP %s port %s", cfg.Address, cfg.Port)
	if err := http.ListenAndServe(cfg.Address+":"+cfg.Port, corsOptions(loggedRouter)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} /*main()*/
