package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sswtshoo/pokedexcli/utils/pokeapi"
)

func commandExplore(c *config, args ...string) error {
	name := args[0]
	if name == "" {
		return errors.New("area name can not be empty")
	} else {
		url := "https://pokeapi.co/api/v2/location-area/" + name
		var areaData pokeapi.LocationAreaDetails
		val, exists := cache.Get(url) 
		if exists {
			if err := json.Unmarshal(val, &areaData); err != nil {
				return errors.New("error parsing cached json data")
			}
		} else {
			data, err := pokeapi.CallAPI(url)
			if err != nil {
				return errors.New("error reading response body")
			}
			cache.Add(url, data)
			if err := json.Unmarshal(data, &areaData); err != nil {
				return errors.New("error parsing json data")
			}
		}
		for _, encounter := range areaData.PokemonEncounters {
			fmt.Println(encounter.Pokemon.Name)
		}
		
	}
	return nil
}