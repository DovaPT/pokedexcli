package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputScanner := bufio.NewScanner(os.Stdin)
	var cfg CommandConfig
	for {
		fmt.Print("pokedex > ")
		inputScanner.Scan()
		input := cleanInput(inputScanner.Text())
		command, exists := cliCommands[input[0]]
		if exists {
			switch input[0] {
			case "explore":
				cfg.Location = &input[1]
			default:
			}
			err := command.callback(&cfg)
			if err != nil {
				fmt.Printf("error: %v", err)
			}
		} else {
			commandHelp(nil)
		}
	}
}
