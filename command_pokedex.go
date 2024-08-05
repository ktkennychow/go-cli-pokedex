package main

import (
	"fmt"
)

func commandPokedex(cfg *config, _ *string, _ *string) error{
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Println("  -", pokemon.Name)
	}
	return nil
}