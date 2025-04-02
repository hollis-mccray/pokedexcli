package main

import "fmt"



func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, command := range listCommands() {
		fmt.Println(command.menuString())
	}
	return nil
}