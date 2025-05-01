package resistorcolor

// Colors returns the list of all colors.

var keys []string 

func init() {
	keys = []string{ "black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white" }
}

func Colors() []string {
	return keys
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	for idx, val := range keys {
		if val == color {
			return idx
		}
	}
	return 0
}
