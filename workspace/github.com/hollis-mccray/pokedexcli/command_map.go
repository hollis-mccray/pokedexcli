package main

import (
	"github.com/hollis-mccray/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	if c.Next == "" {
		fmt.Println("Already at the end of the list.")
		return nil
	}
	locationResp, err := pokeapi.ListLocationData(cfg.Next)
	if err != nil {
		return err
	}

	cfg.Next = locationsResp.Next
	cfg.Previous = locationsResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if c.Previous == "" {
		fmt.Println("Already at the beginning of the list.")
		return nil
	}
	locationResp, err := pokeapi.ListLocationData(c.Previous)
	if err != nil {
		return err
	}

	cfg.Next = locationsResp.Next
	cfg.Previous = locationsResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}