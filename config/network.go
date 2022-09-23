package config

import (
	"log"
	"os"
)

// Dibuat di 1 struct
// Dipanggil di 1 file config (ex. config.go)

func GetBaseUrl() string {
	baseUrl := os.Getenv("BASE_URL_PERSON")
	if baseUrl == "" {
		log.Panic("Cannot get base url")
	}

	return baseUrl
}

func GetApiKey() string {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Panic("Cannot get api key")
	}

	return apiKey
}
