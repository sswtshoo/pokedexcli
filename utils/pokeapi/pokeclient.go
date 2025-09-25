package pokeapi

import (
	"net/http"
	"time"

	"github.com/sswtshoo/pokedexcli/utils/pokecache"
)

type Client struct {
	cache pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, interval time.Duration) Client {
	return Client{
		cache: *pokecache.NewCache(interval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}