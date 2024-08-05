package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, _ *string, pokemonName *string) error{
	fmt.Println("Throwing a Pokeball at " + *pokemonName + "...")

	pokemonDetailsResp, err := cfg.pokeapiClient.GetPokemonDetails(pokemonName)
	if err != nil {
		return err
	}

	pokeballStrength := 50.00
	luck := rand.Float64() 
	difficulty := float64(pokemonDetailsResp.Base_experience)
	catchSuccess := pokeballStrength > luck * difficulty

	if catchSuccess {
		cfg.caughtPokemon[*pokemonName] = pokemonDetailsResp
		fmt.Println(*pokemonName + " was caught!")
		} else {
		fmt.Println(*pokemonName + " escaped!")
	}
	return nil
}