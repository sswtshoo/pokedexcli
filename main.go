package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/sswtshoo/pokedexcli/pokeapi"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

type locationArea struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type responseData struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []locationArea `json:"results"`
}

var nextUrl *string = nil
var prevUrl *string = nil

var currentResponse *responseData = nil

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
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}

		command, exists := commands[scanner.Text()] 
		if exists {
			command.callback()
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	slice := strings.Split(strings.ToLower(trimmed), " ")
	return slice
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
				fmt.Println(`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
			return nil
}

func commandMap() error {
	url := "https://pokeapi.co/api/v2/location-area"
	if nextUrl != nil {
		url = *nextUrl
	}
	var response responseData
	data, err := pokeapi.CallAPI(url)
	if err != nil {
		fmt.Println("Error reading response body")
	}
	if err := json.Unmarshal(data, &response); err != nil {
		fmt.Println("Error parsing JSON data")
	}

	for _, location := range response.Results {
		fmt.Println(location.Name)
	}
	nextUrl = response.Next
	prevUrl = response.Previous
	currentResponse = &response
	return nil
}

func commandMapB() error {
	if prevUrl == nil {
		fmt.Println("you're on the first page")
		results := currentResponse.Results
		for _, location := range results {
		fmt.Println(location.Name)
	}
	} else {
		var response responseData
		data, err := pokeapi.CallAPI(prevUrl)
		if err != nil {
			fmt.Println("Error reading response body")
		}
		if err := json.Unmarshal(data, &response); err != nil {
			fmt.Println("Error parsing JSON data")
		}

		for _, location := range response.Results {
			fmt.Println(location.Name)
		}
		nextUrl = response.Next
		prevUrl = response.Previous
		}
	return nil
}