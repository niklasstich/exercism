package lasagna

// TODO: define the 'OvenTime()' function
func OvenTime() int {
	return 40
}

// TODO: define the 'RemainingOvenTime()' function
func RemainingOvenTime(spentTime int) int {
	return OvenTime() - spentTime
}

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers int) int {
	return 2*layers
}

// TODO: define the 'ElapsedTime()' function
func ElapsedTime(layers, spentTime int) int {
	return PreparationTime(layers) + spentTime
}
