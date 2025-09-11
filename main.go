package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	inputScanner := bufio.NewScanner(os.Stdin)
	for ; ; {
		fmt.Print("pokedex > ")
		inputScanner.Scan()
		input := inputScanner.Text()
		command := cleanInput(input)[0]
		fmt.Printf("Your command was: %s\n", command)
	}
}

func cleanInput(text string) []string{
	words := strings.Fields(text)
	return words
}
