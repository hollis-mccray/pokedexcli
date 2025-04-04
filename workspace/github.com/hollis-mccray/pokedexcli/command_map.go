package main

import (
	"github.com/hollis-mccray/pokedexcli/internal/pokeapi"
)

func commandMap(c Configuration) error {
	pokeapi.GetMaps(c)
	return nil
}