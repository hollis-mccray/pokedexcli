package main

import (
	"fmt"
	"github.com/hollis-mccray/pokedexcli/internal/util"
)

func commandHelp(c configuration) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, command := range listCommands() {
		fmt.Println(command.menuString())
	}
	return nil
}