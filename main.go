package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type WeatherResponse struct {
	City        string `json:"city"`
	Temperature string `json:"temperature"`
	Weather     string `json:"weather"`
}

type CityRequest struct {
	Name string `json:"name"`
}

//need to add get code here will be done by hassans son

func handleCity(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetCity(w, r)
	case http.MethodPost:
		handlePostCity(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetCity(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("name")
	if city == "" {
		http.Error(w, "missing city parameter", http.StatusBadRequest)
		return
	}

	weather, err := getWeather(city)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get weather for %s: %s", city, err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}

func handlePostCity(w http.ResponseWriter, r *http.Request) {
	var cityReq CityRequest
	err := json.NewDecoder(r.Body).Decode(&cityReq)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	city := cityReq.Name
	if city == "" {
		http.Error(w, "missing 'name' field in request body", http.StatusBadRequest)
		return
	}

	weather, err := getWeather(city)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get weather for %s: %s", city, err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}

func main() {
	http.HandleFunc("/city", handleCity)

	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
