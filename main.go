package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	}
}

func main(){
	inputScanner := bufio.NewScanner(os.Stdin)

	cfg := CommandConfig{Next: nil, Prev: nil}

	for ; ; {
		fmt.Print("pokedex > ")
		inputScanner.Scan()
		input := cleanInput(inputScanner.Text())[0]
		command, exists := cliCommands[input]
		if exists {
			err := command.callback(&cfg)
			if err != nil {
				fmt.Printf("error: %v", err)
			}
		} else {
			commandHelp(nil)
		}
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

type LocationData struct{
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		Url string `json:"url"`
	}`json:"results"`
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

func getLocations(path *string) (LocationData, error){
	jsonData, err := queryApi(*path)
	if err != nil {
		return LocationData{}, err
	}
	var locations LocationData
	json.Unmarshal(jsonData, &locations)

	return locations, nil
}


func queryApi(url string) ([]byte, error){
  res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	return jsonData, nil
}
