package main

import (
	"fmt"
	"os"
	"github.com/hollis-mccray/pokedexcli/internal/pokeapi"
)

func commandExit(c *pokeapi.Configuration) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}