package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListlocationList(pageURL string) (locationList, error) {
	url := baseURL + "/location-area"
	if pageURL != "" {
		url = pageURL
	}

	data := locationList{}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := locationList{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return locationList{}, err
		}

		return locationsResp, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return locationList{}, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return locationList{}, err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return locationList{}, err
	}
	return data, nil
}
