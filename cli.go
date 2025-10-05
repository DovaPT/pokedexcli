package main

import (
	"fmt"
	"os"
	"strings"
)

var cliCommands map[string]CliCommand

func init(){
	cliCommands = map[string]CliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays the next 20 locations",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the previous 20 locations",
			callback: commandMapb,
		},
		"explore": {
			name: "explore [location]",
			description: "Shows pokemon found in location.",
			callback: commandExplore,
		},
	}
}


type CliCommand struct {
	name string
	description string
	callback func(*CommandConfig) error
}

type CommandConfig struct {
	Next *string
	Prev *string
	Location *string
}

func cleanInput(text string) []string{
	words := strings.Fields(strings.ToLower(text))
	return words
}

func commandExit(_ *CommandConfig) error{
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *CommandConfig) error {
	println("Welcome to the Pokedex!\nUsage:\n")
	for _, command := range cliCommands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}


func commandMap(cfg *CommandConfig) error {
	if (*cfg).Next == nil {
		next := "https://pokeapi.co/api/v2/location-area/?limit=20"
		cfg.Next = &next
	}
	locationData, err := getLocations(cfg.Next)
	if err != nil {
		return err
	}
	for _, location := range locationData.Results{
		fmt.Printf("%s\n", location.Name)
	}

	cfg.Next = locationData.Next
	cfg.Prev = locationData.Previous
	return nil
}

func commandMapb(cfg *CommandConfig) error {
	if (*cfg).Prev == nil {
		println("You're on the first page")
		return nil
	}
	locationData, err := getLocations(cfg.Next)
	if err != nil {
		return err
	}
	for _, location := range locationData.Results{
		fmt.Printf("%s\n", location.Name)
	}

	cfg.Next = locationData.Next
	cfg.Prev = locationData.Previous
	return nil
}

func commandExplore(cfg *CommandConfig)error {
	if cfg.Location == nil{
		return nil
	}
	pokemon, err := getLocationInfo(cfg.Location)
	if err != nil {
		return err
	}
	fmt.Printf("Exoloring %s...\n", *cfg.Location)
	println("Found Pokemon:")
	for _, mon := range pokemon.PokemonEncounter{
		fmt.Printf(" - %s\n", mon.Pokemon.Name)
	}
	return nil
}
