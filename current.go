package openweathermap

import (
	"encoding/json"
	"fmt"
	"net/url"
)

var apiURL = "http://api.openweathermap.org/data/2.5"

// getForecast sends a constructed request to the OpenWeatherMap API and decodes
// the json response
func getForecast(owm *OpenWeatherMap, url string) (*CurrentWeatherData, error) {
	forecast := &CurrentWeatherData{}
	resp, err := owm.client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = CheckForErrors(resp)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(resp.Body).Decode(forecast); err != nil {
		return nil, &JSONDecodingError{ErrString: "An error occured when decoding response to JSON"}
	}

	return forecast, nil
}

// GetCurrentByCityName retrieves the current weather data for the provided
// location
func (owm *OpenWeatherMap) GetCurrentByCityName(location string) (*CurrentWeatherData, error) {
	if location == "" {
		return nil, &IllegalArgumentError{ErrString: "Location cannot be an empty string"}
	}

	url := fmt.Sprintf("%s/weather?q=%s&units=%s&appid=%s", apiURL, url.QueryEscape(location), owm.unit, owm.key)
	return getForecast(owm, url)
}

// GetCurrentByCoords retrieves the current weather data for the provided
// coordinates
func (owm *OpenWeatherMap) GetCurrentByCoords(coords *Coords) (*CurrentWeatherData, error) {
	url := fmt.Sprintf("%s/weather?lat=%f&lon=%f&units=%s&appid=%s", apiURL, coords.Lat, coords.Lon, owm.unit, owm.key)
	return getForecast(owm, url)
}

// GetCurrentByID retrieves the current weather data for the provided city id
func (owm *OpenWeatherMap) GetCurrentByID(id int) (*CurrentWeatherData, error) {
	url := fmt.Sprintf("%s/weather?id=%d&units=%s&appid=%s", apiURL, id, owm.unit, owm.key)
	return getForecast(owm, url)
}

// GetCurrentByZipCode retrieves the current weather data for the provided zip
// code
func (owm *OpenWeatherMap) GetCurrentByZipCode(zipCode int, countryCode string) (*CurrentWeatherData, error) {
	url := fmt.Sprintf("%s/weather?zip=%d,%s&units=%s&appid=%s", apiURL, zipCode, countryCode, owm.unit, owm.key)
	return getForecast(owm, url)
}
