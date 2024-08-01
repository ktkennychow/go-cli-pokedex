package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ktkennychow/go-cli-pokedex/internal/pokecache"
)

type Config struct {
	Cache *pokecache.Cache
	Next string
	Previous string
}

func startRepl(){
	reader := bufio.NewScanner(os.Stdin)
	cache := pokecache.NewCache(5 * time.Minute)
	cfg := Config{Next: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20", Previous: "", Cache: cache}
	for {
		fmt.Printf("Pokedex > ")
		reader.Scan()
		userInput := reader.Text()
		words := formatInput(userInput)
		commands, commandsWithConfigAndCache := getCommands()
		command, exist := commands[words[0]]
		if !exist {
			command, exist := commandsWithConfigAndCache[words[0]]
			if !exist {
				fmt.Printf("%v is not a valid command.\n", words[0])
				continue
			}
			err := command.callback(&cfg)
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
	callback    func(cfg *Config) error
}

type commands struct {
	name        string
	description string
	callback    func() error
}

func getCommands() (map[string]commands, map[string]commandsWithConfig) {
	commands := map[string]commands{
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
			"mapsf": {
				name:        "mapsf",
				description: "Show the next 20 location areas",
				callback:    commandMapNext,
			},
			"mapsb": {
				name:        "mapsb",
				description: "Show the previous 20 location areas",
				callback:    commandMapPrevious,
			},
		}
	return commands, commandsWithConfig
}

func formatInput(s string) []string{
	s = strings.TrimSpace(s)
	words := strings.Fields(s)
	return words
}