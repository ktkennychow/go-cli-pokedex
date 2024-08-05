package main

import (
	"fmt"
)

func commandInspect(cfg *config, _ *string, pokemonName *string) error{
	pokemon, exist:= cfg.caughtPokemon[*pokemonName]
	if !exist {
		fmt.Println("You have not caught that pokemon")
		return nil
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats: ")
	for _, stat := range pokemon.Stats {
		fmt.Println("  -", stat.Stat.Name + ":", stat.Base_stat)
	}
	fmt.Println("Types: ")
	for _, pokeType := range pokemon.Types {
		fmt.Println("  -", pokeType.Type.Name)
	}

	return nil
}