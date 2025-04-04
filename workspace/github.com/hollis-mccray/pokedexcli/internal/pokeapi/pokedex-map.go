package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

type locationData struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetMaps(c configuration) error {
	addr := *c.Next
	res, err := http.Get(addr)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	data := locationData{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}
	for _, location := range data.Results {
		fmt.Println(location.Name)
	}
	return nil
}