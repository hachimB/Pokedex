package main

import (
	"testing"
	"github.com/hachimB/Pokedex/internal/repl"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  hello  world  ",
			expected: []string{"hello", "world"}, 
		},
		{
			input: "Hello Everyone", 
			expected: []string{"hello", "everyone"},
		},
		{
			input: "Hi, How are you ?",
			expected: []string{"hi,", "how", "are", "you", "?"},
		},
	}

	for _, c := range cases {
	actual := repl.CleanInput(c.input)
	if len(actual) != len(c.expected) {
		t.Errorf("Error⚠️: You're missing some words!")
		t.Fail()
	} 
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		if word != expectedWord {
			t.Errorf("Error⚠️: Different words!")
			t.Fail()
		}
		}
	}
}