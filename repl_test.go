package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  what is that pokemon  ",
			expected: []string{"what", "is", "that", "pokemon"},
		},
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    "one",
			expected: []string{"one"},
		},
		// add more cases here
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		if len(actual) != len(c.expected) {
			t.Errorf("Actual len: %d, Expected: %d", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice.
			if word != expectedWord {
				t.Errorf("Actual word: %s, Expected Word: %s", word, expectedWord)
			}
		}
	}
}
