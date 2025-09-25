package main

import (
	"fmt"
)

func commandPokedex(c *config, args ...string) error {
	fmt.Print("Your Pokedex\n")
	for _, pokemon := range c.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}