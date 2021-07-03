package reverse

func Reverse(s string) string {
	runes := []rune(s)
	// Have two values run towards each other
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		// Swap the values at the indices
		runes[i], runes[j] = runes[j], runes[i]
	}
	// Turn slice of runes back into a string
	return string(runes)
}
