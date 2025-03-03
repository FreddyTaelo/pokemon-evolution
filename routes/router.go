package routes

import (
	"net/http"
	"pokemon-evolution/controllers"
	"pokemon-evolution/middleware"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/evolution/{id:[0-9]+}", controllers.GetEvolutionChain).Methods("GET")

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Healthy"))
	}).Methods("GET")

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	router.Use(middleware.LoggingMiddleware)

	return router
} /* RegisterRoutes()*/
