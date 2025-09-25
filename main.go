package main

import (
	"time"

	"github.com/sswtshoo/pokedexcli/utils/pokeapi"
	"github.com/sswtshoo/pokedexcli/utils/pokecache"
)

// var nextUrl *string = nil
// var prevUrl *string = nil

var cache *pokecache.Cache = pokecache.NewCache(time.Millisecond * 5000)

// var currentResponse *pokeapi.ResponseData = nil

// var caughtPokemon = make(map[string]pokeapi.CatchPokemon)



// var argName *string




func main() {
	pokeclient := pokeapi.NewClient(5000 * time.Millisecond, 5 * time.Minute)

	c := &config{
		caughtPokemon: map[string]pokeapi.CatchPokemon{},
		pokeapiclient: pokeclient,
	}

	startREPL(c)
}