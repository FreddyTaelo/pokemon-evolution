// config/config.go
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PokeAPIURL string
	Port       string
	Address    string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default environment variables.")
	}

	return &Config{
		PokeAPIURL: getEnv("POKEAPI_URL", "https://pokeapi.co/api/v2"),
		Port:       getEnv("PORT", "5001"),
		Address:    getEnv("ADDR", "0.0.0.0"),
	}
} /*LoadConfig()*/

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
} /*getEnv()*/
