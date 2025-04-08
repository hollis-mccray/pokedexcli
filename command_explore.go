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
	fmt.Printf("Exploring %s...\n", location)
	encounters, err := cfg.pokeapiClient.ListEncounters(location)
	if err != nil {
		return err
	}
	if len(encounters) != 0 {
		for _, pokemon := range encounters {
			fmt.Println(pokemon)
		}
	} else {
		fmt.Println("No Pokemon found.")
	}
	return nil
}
