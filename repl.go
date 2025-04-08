package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hollis-mccray/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	Next          string
	Previous      string
	Arguments     []string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
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
		if len(words) >= 2 {
			cfg.Arguments = words[1:]
		} else {
			cfg.Arguments = []string{}
		}
		err := command.callback(cfg)
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
		"mapb": {
			name:        "mapb",
			description: "Displays the previous twenty locations in Pokemon",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays Pokemon found in a location",
			callback:    commandExplore,
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
