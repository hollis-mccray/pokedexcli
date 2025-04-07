package pokeapi

import (
	"io"
	"net/http"
	"encoding/json"
)

func ListlocationList(pageURL *string) (locationList, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	data := locationList{}
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