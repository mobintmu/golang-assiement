package main

import (
	"aqary-2/rearrange"
	"fmt"
	"testing"
)

func TestCasesTest(t *testing.T) {

	t.Run("Test Case 1", func(t *testing.T) {

		input := "aab"
		expected := "aba"

		output := rearrange.Rearrange(input)

		if len(output) != len(expected) {
			t.Errorf("Test Case 1 failed: %v inputted, %v expected, recieved: %v", input, expected, output)
		}

	})

	t.Run("Test Case 2", func(t *testing.T) {

		input := "aaab"
		expected := ""
		output := rearrange.Rearrange(input)

		if len(output) != len(expected) {

			t.Errorf("Test Case 2 failed: %v inputted, %v expected, recieved: %v", input, expected, output)
		}

	})

	t.Run("Test Case 3", func(t *testing.T) {

		input := "aaabbbc"
		expected := "abababc"
		output := rearrange.Rearrange(input)

		if len(output) != len(expected) {
			t.Errorf("Test Case 3 failed: %v inputted, %v expected, recieved: %v", input, expected, output)
		}

	})

	t.Run("Test Case 4", func(t *testing.T) {

		input := "aaabbbcc"
		output := rearrange.Rearrange(input)
		expected := "abababcc"

		if len(output) != len(expected) {
			t.Errorf("Test Case 4 failed: %v inputted, %v expected, recieved: %v", input, expected, output)
		}

	})

	t.Run("Test Case 5", func(t *testing.T) {

		input := "aaabbbccc"
		expected := "abcabcabc"

		output := rearrange.Rearrange(input)

		if len(output) != len(expected) {

			t.Errorf("Test Case 5 failed: %v inputted, %v expected, recieved: %v", input, expected, output)
		}

	})

	t.Run("Test Case 6", func(t *testing.T) {

		input := "aaabbbcccc"
		expected := "cabcabcabc"
		output := rearrange.Rearrange(input)

		fmt.Println(input, output)

		if output != expected {

			t.Errorf("Test Case 6 failed: %v inputted, %v expected, recieved: %v", input, expected, output)
		}

	})

	t.Run("Test Case 7", func(t *testing.T) {

		input := "aaab"
		output := rearrange.Rearrange(input)
		expected := ""

		if len(output) != len(expected) {

			t.Errorf("Test Case 7 failed: %v inputted, %v expected, recieved: %v", input, expected, output)
		}
	})
}
