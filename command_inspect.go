package main

import (
	"fmt"
)

func commandInspect (c *config, args ...string) error {
	name := args[0]
	pokemon, exists := c.caughtPokemon[name]
	if !exists {
		fmt.Printf("%s hasn't been caught", name)
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stats := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stats.Stat.Name, stats.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, types := range pokemon.Types {
		fmt.Printf("  - %s\n", types.Type.Name)
	}
	return nil
}