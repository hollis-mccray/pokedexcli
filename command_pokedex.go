package main

import "fmt"

func commandPokedex(cfg *config, args []string) error {
	fmt.Println("Your Pokedex:")
	for name := range cfg.pokeapiClient.Pokedex {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
