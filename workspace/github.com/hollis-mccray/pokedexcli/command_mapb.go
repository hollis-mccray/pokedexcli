package main

import (
	"github.com/hollis-mccray/pokedexcli/internal/pokeapi"
)

func commandMapb(c *pokeapi.Configuration) error {
	pokeapi.Mapb(c)
	return nil
}