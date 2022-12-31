// Package weather for Goblinocus forecast.
package weather

// CurrentCondition represent the current weather condition for the Forecast Function.
var CurrentCondition string

// CurrentLocation represent the current weather location for the Forecast Function.
var CurrentLocation string

// Forecast represent the weather condition of the given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
