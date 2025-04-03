package main

import "github.com/hollis-mccray/pokedex/internal/pokedex-api"

func commandMap(c cliConfig) error {
	pokedex-api.getMaps(c)
	return nil
}