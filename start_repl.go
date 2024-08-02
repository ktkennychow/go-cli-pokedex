package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ktkennychow/go-cli-pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
}

func startRepl(cfg *config){
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		reader.Scan()

		userInput := reader.Text()
		words := formatInput(userInput)
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		commands:= getCommands()
		command, exist := commands[commandName]
		if exist {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("%v is not a valid command.\n", words[0])
			continue
		}
	}
}

type commands struct {
	name        string
	description string
	callback    func(cfg *config) error
}

func getCommands() (map[string]commands) {
	return map[string]commands{
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
			"mapf": {
				name:        "mapf",
				description: "Show the next 20 location areas",
				callback:    commandMapNext,
			},
			"mapb": {
				name:        "mapb",
				description: "Show the previous 20 location areas",
				callback:    commandMapPrevious,
			},
		}
}

func formatInput(s string) []string{
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	words := strings.Fields(s)
	return words
}