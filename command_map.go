package main

import (
	"errors"
	"fmt"

	"github.com/ktkennychow/go-cli-pokedex/internal/pokeapi"
)

func commandMapNext(c *Config) error{
	newRes, err := pokeapi.GetPokedexApi(c.Next)
	if err != nil {
		return err
	}
	
	for _, locationArea := range newRes.LocationAreas{
		fmt.Println(locationArea.Name)
	}
	c.Next, c.Previous = newRes.Next, newRes.Previous
	return nil
}

func commandMapPrevious(c *Config) error{
	if (c.Previous == "") {
		return errors.New("already on the first page")
	}

	newRes, err := pokeapi.GetPokedexApi(c.Previous)
	if err != nil {
		return err
	}
	
	for _, locationArea := range newRes.LocationAreas{
		fmt.Println(locationArea.Name)
	}
	c.Next, c.Previous = newRes.Next, newRes.Previous
	return nil
}