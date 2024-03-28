package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
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

func getWeather(city string) (*WeatherResponse, error) {
	apiKey := "d51319b8aafa1e0618c55136562d617b"
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	client := resty.New()
	resp, err := client.R().Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch weather data: %s", resp.Status())
	}

	var weatherData map[string]interface{}
	err = json.Unmarshal(resp.Body(), &weatherData)
	if err != nil {
		return nil, err
	}

	weather := WeatherResponse{
		City:        weatherData["name"].(string),
		Temperature: fmt.Sprintf("%.1f°C", weatherData["main"].(map[string]interface{})["temp"].(float64)),
		Weather:     weatherData["weather"].([]interface{})[0].(map[string]interface{})["main"].(string),
	}

	return &weather, nil
}

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
