package main

import (
	"github.com/hollis-mccray/pokedexcli/internal/pokeapi"
)

func commandMap(c *pokeapi.Configuration) error {
	pokeapi.Map(c)
	return nil
}