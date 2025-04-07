package main

import (
	"fmt"
)

func commandExplore(cfg *config) error {
	if len(cfg.Arguments) == 0 {
		fmt.Println("Location not specified")
		return nil
	}
	location := cfg.Arguments[0]
	fmt.Printf("Exploring %s...", location)
	encounters := pokeapi.ListEncounters(location)
	return nil
}