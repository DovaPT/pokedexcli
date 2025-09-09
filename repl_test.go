package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello   world   ",
			expected: []string{"hello", "world"},
		},
		{
			input: "  help       me           friend.",
			expected: []string{"help", "me", "friend."},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) < len(c.expected){
			t.Error("actual has less values than expected")
			t.Fail()
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			fmt.Printf("actual: %s \nexpected: %s", actual, expectedWord)
			if word != expectedWord {
				t.Errorf("actual: %s \nexpected: %s", word, expectedWord)
				t.Fail()
			}
		}
	}
}
