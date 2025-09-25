package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sswtshoo/pokedexcli/utils/pokeapi"
)

func commandMap(c *config, args ...string) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if c.nextLocationUrl != nil {
		url = *c.nextLocationUrl
	}
	var response pokeapi.ResponseData
	val, exists := cache.Get(url)
	if exists {
		if err := json.Unmarshal(val, &response); err != nil {
			fmt.Println("Error parsing cached JSON data")
		}
	} else {
		data, err := pokeapi.CallAPI(url)
		if err != nil {
			fmt.Println("Error reading response body")
		}
		cache.Add(url, data)
		if err := json.Unmarshal(data, &response); err != nil {
			fmt.Println("Error parsing JSON data")
		}
	}
	for _, location := range response.Results {
		fmt.Println(location.Name)
	}
	c.nextLocationUrl = response.Next
	c.previousLocationUrl = response.Previous

	return nil
}

func commandMapB(c *config, args ...string) error {
	if c.previousLocationUrl == nil {
		return errors.New("you're on the first page")
	} else {
		val, exists := cache.Get(*c.previousLocationUrl)
		var response pokeapi.ResponseData
		if exists {
			if err := json.Unmarshal(val, &response); err != nil {
				fmt.Println("Error parsing cached JSON data")
			}
		} else {
			data, err := pokeapi.CallAPI(*c.previousLocationUrl)
			if err != nil {
				fmt.Println("Error reading response body")
			}
			if err := json.Unmarshal(data, &response); err != nil {
				fmt.Println("Error parsing JSON data")
			}
			cache.Add(*c.previousLocationUrl, data)
		}

		for _, location := range response.Results {
			fmt.Println(location.Name)
		}
		c.nextLocationUrl = response.Next
		c.previousLocationUrl = response.Previous
		}
	return nil
}