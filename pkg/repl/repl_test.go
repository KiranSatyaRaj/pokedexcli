package repl

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "    hello   world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "   Charmander Bulbasaur        PIKACHU         ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("expected %v, got %v", c.expected, actual)
		}

		for i := 0; i < len(actual); i++ {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected %s, got %s", expectedWord, word)
			}
		}
	}
}

