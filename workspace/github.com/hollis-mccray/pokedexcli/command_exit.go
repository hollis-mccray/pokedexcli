package main

import (
	"fmt"
	"os"
	"github.com/hollis-mccray/pokedexcli/internal/util"
)

func commandExit(c configuration) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}