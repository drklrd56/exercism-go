// Package weather provides a function to get the current weather forecast based on the CurrentCondition and the CurrentLocation.
package weather

// CurrentCondition holds the current weather condition.
var CurrentCondition string
// CurrentLocation holds the current location.
var CurrentLocation string

// Forecast returns the current weather forecast based on the city and the condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
