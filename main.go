package main

import (
	"fmt"
	"strings"
)

func main(){
	var actual, expected string
	expected = "test"

	if actual != expected {
		println("Success")
	} else {
		println("Failed")
	}
}

func cleanInput(text string) []string{
	words := strings.Fields(text)
	fmt.Printf("%v\n", words)
	return words
}
