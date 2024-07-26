package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(){
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		reader.Scan()
		userInput := reader.Text()
		commands := formatInput(userInput)
		v, exist := getCommands()[commands[0]]
		if !exist {
			fmt.Printf("%v is not a valid command.\n", commands[0])
		} else {
			err := v.callback()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
}

func formatInput(s string) []string{
	s = strings.TrimSpace(s)
	words := strings.Fields(s)
	return words
}