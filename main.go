package main

import (
	"bufio"
	"fmt"
	"os"
)

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
