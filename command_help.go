package main

import (
	"fmt"
)

func commandHelp(_ *config, _ *string, _ *string) error {
	commands:= getCommands()
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, v := range commands {
		fmt.Printf("%v: %v\n", v.name, v.description)
	} 
	fmt.Println("")
	return nil
}