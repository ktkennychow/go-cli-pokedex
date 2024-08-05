package main

import "fmt"

func commandExplore(cfg *config, areaName *string, _ *string) error{
	explorelocationResp, err := cfg.pokeapiClient.ExploreLocationArea(areaName)
	if err != nil {
		return err
	}

	fmt.Println("Exploring " + *areaName + "...")
	fmt.Println("Found Pokemon:")
	for _, encounter := range explorelocationResp.Pokemon_encounters{
			fmt.Println(" - ", encounter.Pokemon.Name)
	}
	return nil
}