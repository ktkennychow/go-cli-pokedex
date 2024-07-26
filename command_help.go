package main

import (
	"fmt"
)

func commandHelp() error {
	cliCommands := getCommands()
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, v := range cliCommands {
		fmt.Printf("%v: %v\n", v.name, v.description)
	} 
	fmt.Println("")
	return nil
}