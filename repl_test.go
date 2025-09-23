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
			input: " hello world",
			expected: []string{"hello", "world"},
		}, 
		{
			input: "Mitochondria is the powerhouse of the cell!",
			expected: []string{"mitochondria", "is", "the", "powerhouse", "of", "the", "cell!"},
		},
		{
			input: "We are going to the moon!",
			expected: []string{"we", "are", "going", "to", "the", "moon!"},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("Failed! \nExpecting: %v, Got: %v", cs.expected, actual)
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


