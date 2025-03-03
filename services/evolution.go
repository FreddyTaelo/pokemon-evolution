// services/evolution.go
package services

import (
	"encoding/json"
	"fmt"
	"pokemon-evolution/api"
	"pokemon-evolution/config"
	"pokemon-evolution/models"
)

var cfg = config.LoadConfig()

// only works on one level
func GetEvolutionChain(id int) (*models.Evolution, error) {
	url := fmt.Sprintf("%s/evolution-chain/%d", cfg.PokeAPIURL, id)
	resp, err := api.GetClient().Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Chain struct {
			Species struct {
				Name string `json:"name"`
			} `json:"species"`
			EvolvesTo []struct {
				Species struct {
					Name string `json:"name"`
				} `json:"species"`
				EvolutionDetails []struct {
					Trigger struct {
						Name string `json:"name"`
					} `json:"trigger"`
					MinLevel int `json:"min_level"`
				} `json:"evolution_details"`
			} `json:"evolves_to"`
		} `json:"chain"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	evolution := &models.Evolution{
		Name:      data.Chain.Species.Name,
		EvolvesTo: []models.Evolution{},
	}

	for _, e := range data.Chain.EvolvesTo {
		evolution.EvolvesTo = append(evolution.EvolvesTo, models.Evolution{
			Name: e.Species.Name,
			Details: []models.EvolutionDetail{
				{
					Trigger:  e.EvolutionDetails[0].Trigger.Name,
					MinLevel: e.EvolutionDetails[0].MinLevel,
				},
			},
		})
	}

	return evolution, nil
}
