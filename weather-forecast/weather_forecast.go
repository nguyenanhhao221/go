// Package weather for Goblinocus forecast.
package weather

var CurrentCondition string // CurrentCondition represent the current weather condition for the Forecast Function.
var CurrentLocation string

func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
