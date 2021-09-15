//Package weather provides methods for formatting current weather circumstances for logging purposes.
package weather

var (
	//CurrentCondition stores the current weather conditions in a given location.
	CurrentCondition string
	//CurrentLocation stores the current location regarding CurrentCondition.
	CurrentLocation  string
)

//Log formats a given city and weather condition into pretty-print for logging purposes.
func Log(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
