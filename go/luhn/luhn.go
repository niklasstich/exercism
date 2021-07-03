package main

import (
	"strconv"
	"strings"
)

func Valid(input string) bool {
	// Remove all whitespace
	input = strings.Join(strings.Fields(input), "")
	// Single character input is invalid
	if len(input) == 1 {
		return false
	}
	// Split every letter of the string into a slice
	inputArray := strings.Split(input, "")
	// Reverse the slice so it's easier to pick the numbers later
	for i, j := 0, len(inputArray)-1; i < j; i, j = i+1, j-1 {
		inputArray[i], inputArray[j] = inputArray[j], inputArray[i]
	}
	sum := 0
	// Iterate over every character in the slice and convert into an integer
	for i := 0; i < len(inputArray); i += 1 {
		val, err := strconv.Atoi(inputArray[i])
		if err != nil {
			return false
		}
		// If mod 2 is uneven, we have to double the character and sub 9 if over 9
		if i%2 == 1 {
			val = val * 2
			if val > 9 {
				val = val - 9
			}
		}
		// Add the value to the sum
		sum += val
	}
	if sum%10 == 0 {
		return true
	}
	return false
}
