// Package weather provides tools to print out current weather for certain location.
package weather

// CurrentCondition variable represents the current condition/weather.
var CurrentCondition string

// CurrentLocation variable represents the current location.
var CurrentLocation string

// Forecast function prints out the current location with condition based on input.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
