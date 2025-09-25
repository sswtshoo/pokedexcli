package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"

	"github.com/sswtshoo/pokedexcli/utils/pokeapi"
)

func commandCatch(c *config, args ...string) error {
	name := args[0]
	if name == "" {
		return errors.New("pokemon name can not be empty")
	} else {
		fmt.Printf("Throwing a Pokeball at %s...\n", name)
		url := "https://pokeapi.co/api/v2/pokemon/" + name
		var pokemon pokeapi.CatchPokemon
		data, err := pokeapi.CallAPI(url)
		if err != nil {
			return errors.New("error reading response body")
		}
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return errors.New("error parsing json data")
		}

		randomNum := rand.Intn(pokemon.BaseExperience)
		if randomNum > 40 {
			fmt.Printf("%s escaped\n", name)
			return nil
		} 
		fmt.Printf("%s was caught\n", name)
		fmt.Printf("You can now see %s's stats using inspect command\n", name)
		c.caughtPokemon[name] = pokemon
	}
	return nil
}