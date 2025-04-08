package main

import (
	"fmt"
)

func commandHelp(cfg *config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range listCommands() {
		fmt.Println(command.menuString())
	}
	return nil
}
