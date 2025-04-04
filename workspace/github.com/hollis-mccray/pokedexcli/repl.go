package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/hollis-mccray/pokedexcli/internal/util"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	n := "https://pokeapi.co/api/v2/location-area/"
	config.Previous = nil
	config.Next = &n
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())

		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		command, ok := listCommands()[commandName]
		if !ok {
			fmt.Println("Unknown Command")
			continue
		}
		err := command.callback(config)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(configuration) error
}

func (c cliCommand) menuString() string {
	return fmt.Sprintf("%s: %s", c.name, c.description)
}

func listCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Displays the next twenty locations in Pokemon",
			callback:    commandMap,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}