package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"weather-cli/models"
	"weather-cli/pkg/helpers"
)

const apiKey = "7b49eae9ef3d4a31bfa171753231203&q"
const apiUrl = "https://api.weatherapi.com/v1/"

func GetWeather(city string) (models.ForecastDay, error) {
	res, err := http.Get(apiUrl + "forecast.json?key=" + apiKey + "=" + city)

	if err != nil {
		fmt.Printf("Request failed: %v", err)

		return models.ForecastDay{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("Reading body failed: %v", err)

		return models.ForecastDay{}, err
	}

	bodyString := string(body)

	if strings.Contains(bodyString, "error") {
		return models.ForecastDay{}, fmt.Errorf(helpers.FirstToLower("error"))
	}

	forecast := models.ForecastDay{}

	err = json.Unmarshal([]byte(bodyString), &forecast)

	if err != nil {
		fmt.Printf("Unmarshaling failed: %v", err)
	}

	return forecast, nil
}
