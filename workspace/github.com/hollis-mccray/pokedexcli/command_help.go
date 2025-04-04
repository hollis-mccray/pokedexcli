package main

import (
	"fmt"
	"github.com/hollis-mccray/pokedexcli/internal/pokeapi"
)

func commandHelp(c configAPI) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, command := range listCommands() {
		fmt.Println(command.menuString())
	}
	return nil
}