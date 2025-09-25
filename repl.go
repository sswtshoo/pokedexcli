package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sswtshoo/pokedexcli/utils/pokeapi"
)

var commands = map[string]cliCommand{
	"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},
	"help": {
		name: "help",
		description: "Pokedex REPL usage",
		callback: commandHelp,
	},
	"map": {
		name: "map",
		description: "Displays 20 location areas in the Pokemon world",
		callback: commandMap,
	},
	"mapb": {
		name: "mapb",
		description: "Display the previous page of location areas",
		callback: commandMapB,
	},
	"explore": {
		name: "explore <location_name>",
		description: "Display a list of all Pokemon in the given are",
		callback: commandExplore,
	},
	"catch": {
		name: "catch <pokemon>",
		description: "Adds a pokemon to the user's pokedex",
		callback: commandCatch,
	},
	"inspect": {
		name: "inspect <pokemon>",
		description: "Display the stats of a pokemon caught by the user",
		callback: commandInspect,
	},
	"pokedex": {
		name: "pokedex",
		description: "Display the names of all pokemon caught by the user",
		callback: commandPokedex,
	},
}

type config struct{
	pokeapiclient pokeapi.Client
	nextLocationUrl *string
	previousLocationUrl *string
	caughtPokemon map[string]pokeapi.CatchPokemon
}

type cliCommand struct {
	name string
	description string
	callback func(*config, ...string) error
}

func startREPL(c *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			continue
		}

		cleanedInput := cleanInput((scanner.Text()))
		commandName := cleanedInput[0]
		args := []string{}
		if len(cleanedInput) > 1 {
			args = cleanedInput[1:]
		}

		command, exists := commands[commandName]
		if exists {
			err := command.callback(c, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("unknown command")
			continue
		}

	}
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	slice := strings.Split(strings.ToLower(trimmed), " ")
	return slice
}


