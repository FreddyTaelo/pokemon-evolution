package controllers

import (
	"encoding/json"
	"net/http"
	"pokemon-evolution/services"
	"strconv"

	"github.com/gorilla/mux"
)

// @Summary Get Evolution Chain
// @Description Get evolution details for a Pokémon by evolution chain ID.
// @Tags Evolution
// @Accept  json
// @Produce  json
// @Param id path int true "Evolution Chain ID"
// @Success 200 {object} models.Evolution
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/evolution/{id} [get]
func GetEvolutionChain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	evolution, err := services.GetEvolutionChain(id)
	if err != nil {
		http.Error(w, "Error fetching evolution chain", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(evolution)
} /* GetEvolutionChain() */
