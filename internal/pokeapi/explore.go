package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(location string) (LocationDetails, error) {
	url := baseURL + "/location-area/" + location
	if location == "" {
		return LocationDetails{}, nil
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := LocationDetails{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationDetails{}, err
		}

		return locationsResp, nil
	}

	data := LocationDetails{}
	res, err := http.Get(url)
	if err != nil {
		return LocationDetails{}, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationDetails{}, err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return LocationDetails{}, err
	}

	return data, nil
}
