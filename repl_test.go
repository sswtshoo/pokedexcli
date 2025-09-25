package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		}, 
		{
			input: "Lessa GO",
			expected: []string{"lessa", "go"},
		},
		{
			input: "i wOnDer",
			expected: []string{"i", "wonder"},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("Mismatched lengths, expected: %d, got: %d", len(cs.expected), len(actual))
		}
		flag := true
		for i := range actual {
			word := actual[i]
			if word != cs.expected[i] {
				flag = false
			}
		}
		if !flag {
			t.Errorf("Failed! \nExpecting: %v, Got: %v", cs.expected, actual)
		}
	}
}
