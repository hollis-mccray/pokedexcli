package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListEncounters(location string) ([]string, error) {
	url := baseURL + "/location-area/" + location
	if location == "" {
		return []string{}, nil
	}

	data := locationDetails{}
	res, err := http.Get(url)
	if err != nil {
		return []string{}, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []string{}, err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return []string{}, err
	}
	encounters := []string{}

	for _, encounter := range data.PokemonEncounters {
		encounters = append(encounters, encounter.Pokemon.Name)
	}

	return encounters, nil
}
