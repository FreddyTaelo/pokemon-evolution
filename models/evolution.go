// models/evolution.go
package models

type EvolutionDetail struct {
	Trigger   string `json:"trigger"`
	MinLevel  int    `json:"min_level,omitempty"`
	Item      string `json:"item,omitempty"`
	Happiness int    `json:"happiness,omitempty"`
	KnownMove string `json:"known_move,omitempty"`
}

type Evolution struct {
	ID        int               `json:"id"`
	Name      string            `json:"name"`
	EvolvesTo []Evolution       `json:"evolves_to"`
	Details   []EvolutionDetail `json:"details,omitempty"`
}
