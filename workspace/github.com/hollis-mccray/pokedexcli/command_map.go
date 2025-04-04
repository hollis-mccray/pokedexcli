package main

import (
	"github.com/hollis-mccray/pokedexcli/internal/util"
	"github.com/hollis-mccray/pokedexcli/internal/pokeapi"
)

func commandMap(c configuration) error {
	pokeapi.GetMaps(c)
	return nil
}