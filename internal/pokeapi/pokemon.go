package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemon string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon
	if pokemon == "" {
		return Pokemon{}, errors.New("You must provide a Pokemon name")
	}

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemonResp, nil
	}

	data := Pokemon{}
	res, err := http.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return Pokemon{}, err
	}

	return data, nil
}
