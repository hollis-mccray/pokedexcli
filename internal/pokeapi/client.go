package pokeapi

import (
	"net/http"
	"time"

	"github.com/hollis-mccray/pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
	Pokedex    map[string]Pokemon
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
