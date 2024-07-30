package main

import (
	"fmt"
)

func commandHelp() error {
	commandsWithoutConfig, commandsWithConfig := getCommands()
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, v := range commandsWithoutConfig {
		fmt.Printf("%v: %v\n", v.name, v.description)
	} 
	for _, v := range commandsWithConfig {
		fmt.Printf("%v: %v\n", v.name, v.description)
	} 
	fmt.Println("")
	return nil
}