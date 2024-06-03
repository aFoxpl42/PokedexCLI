package main

import (
	"time"

	pokeapi "github.com/aFoxpl42/PokedexCLI/internal/pokeAPI"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}
	startREPL(&cfg)
}