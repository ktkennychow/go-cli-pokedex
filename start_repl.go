package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Next string
	Previous string
}

func startRepl(){
	reader := bufio.NewScanner(os.Stdin)
	c := Config{Next: "https://pokeapi.co/api/v2/location-area", Previous: ""}
	for {
		fmt.Printf("Pokedex > ")
		reader.Scan()
		userInput := reader.Text()
		words := formatInput(userInput)
		commandsWithoutConfig, commandsWithConfig := getCommands()
		command, exist := commandsWithoutConfig[words[0]]
		if !exist {
			command, exist := commandsWithConfig[words[0]]
			if !exist {
				fmt.Printf("%v is not a valid command.\n", words[0])
			}
			err := command.callback(&c)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

type commandsWithConfig struct {
	name        string
	description string
	callback    func(c *Config) error
}

type commandsWithoutConfig struct {
	name        string
	description string
	callback    func() error
}

func getCommands() (map[string]commandsWithoutConfig, map[string]commandsWithConfig) {
	commandsWithoutConfig := map[string]commandsWithoutConfig{
			"help": {
				name:        "help",
				description: "Displays a help message",
				callback:    commandHelp,
			},
			"exit": {
				name:        "exit",
				description: "Exit the Pokedex",
				callback:    commandExit,
			},
		}
	commandsWithConfig := map[string]commandsWithConfig{
			"map": {
				name:        "mapsf",
				description: "Show the next 20 location areas",
				callback:    commandMapNext,
			},
			"mapb": {
				name:        "mapsb",
				description: "Show the previous 20 location areas",
				callback:    commandMapPrevious,
			},
		}
	return commandsWithoutConfig, commandsWithConfig
}

func formatInput(s string) []string{
	s = strings.TrimSpace(s)
	words := strings.Fields(s)
	return words
}