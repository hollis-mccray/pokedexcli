package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("You must provide a Pokemon name")
	}
	name := args[0]

	info, ok := cfg.pokeapiClient.Pokedex[name]

	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", info.Name)
	fmt.Printf("Height: %d\n", info.Height)
	fmt.Printf("Weight: %d\n", info.Weight)
	fmt.Println("Stats:")
	for _, stat := range info.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, stat := range info.Types {
		fmt.Printf("  - %s\n", stat.Type.Name)
	}
	return nil
}
