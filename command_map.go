package main

import (
	"errors"
	"fmt"
)

func commandMapNext(cfg *config, areaName *string, pokemonName *string) error{
	locationAreasResp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationAreasResp.Next
	cfg.prevLocationURL = locationAreasResp.Previous

	for _, locationArea := range locationAreasResp.Results{
		fmt.Println(locationArea.Name)
	}
	return nil
}

func commandMapPrevious(cfg *config, areaName *string, _ *string) error{
	if cfg.prevLocationURL == nil {
		return errors.New("you're on the first page")
	}
	
	locationAreasResp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationAreasResp.Next
	cfg.prevLocationURL = locationAreasResp.Previous

	for _, locationArea := range locationAreasResp.Results{
		fmt.Println(locationArea.Name)
	}
	return nil
}