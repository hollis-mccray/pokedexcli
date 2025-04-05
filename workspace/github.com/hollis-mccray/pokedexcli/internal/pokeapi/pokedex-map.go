package pokeapi

import (
	"io"
	"net/http"
	"encoding/json"
)

type locationData struct {
	Count    int     `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func ListLocationData(url string) (locationData, error) {
	data := locationData{}
	res, err := http.Get(url)
	if err != nil {
		return locationData{}, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return locationData{}, err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return locationData{}, err
	}
	return data, nil
}