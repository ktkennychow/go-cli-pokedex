package main

import (
	"errors"
	"fmt"

	"github.com/ktkennychow/go-cli-pokedex/internal/pokeapi"
)

func commandMapNext(cfg *Config) error{
	newRes, err := pokeapi.GetPokedexApi(cfg.Next, cfg.Cache)
	if err != nil {
		return err
	}
	for _, locationArea := range newRes.LocationAreas{
		fmt.Println(locationArea.Name)
	}
	cfg.Next, cfg.Previous = newRes.Next, newRes.Previous
	return nil
}

func commandMapPrevious(cfg *Config) error{
	if (cfg.Previous == "") {
		return errors.New("already on the first page")
	}

	newRes, err := pokeapi.GetPokedexApi(cfg.Previous, cfg.Cache)
	if err != nil {
		return err
	}
	
	for _, locationArea := range newRes.LocationAreas{
		fmt.Println(locationArea.Name)
	}
	cfg.Next, cfg.Previous = newRes.Next, newRes.Previous
	return nil
}