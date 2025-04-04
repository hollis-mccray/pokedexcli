package main

import (
	"github.com/hollis-mccray/pokedexcli/internal/pokeapi"
	"fmt"
)

func commandMap(c *pokeapi.Configuration) error {
	pokeapi.GetMaps(c)
	return nil
}