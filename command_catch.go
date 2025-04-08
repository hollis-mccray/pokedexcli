package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("You must provide a Pokemon name")
	}
	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	xp := pokemon.BaseExperience
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	chance := rand.Intn(1000)
	if chance <= xp {
		fmt.Printf("%s was caught!\n", name)
		cfg.pokeapiClient.Pokedex[name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}
