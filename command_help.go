package main

import "fmt"


func commandHelp(c *config, args ...string) error {
				fmt.Println(`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
map: Displays a list of location areas, each call fetches the next 20 results
mapb: Display the previous page of location areas
explore <location_area>: display available pokemon in an area
catch <pokemon_name>: attempts to catch a pokemon
inspect <pokemon_name>: Displays properties of pokemon caught by the user`)
			return nil
}