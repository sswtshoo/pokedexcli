package main

import (
	"time"

	"github.com/sswtshoo/pokedexcli/utils/pokeapi"
	"github.com/sswtshoo/pokedexcli/utils/pokecache"
)

var cache *pokecache.Cache = pokecache.NewCache(time.Millisecond * 5000)

func main() {
	pokeclient := pokeapi.NewClient(5000 * time.Millisecond, 5 * time.Minute)

	c := &config{
		caughtPokemon: map[string]pokeapi.CatchPokemon{},
		pokeapiclient: pokeclient,
	}

	startREPL(c)
}
