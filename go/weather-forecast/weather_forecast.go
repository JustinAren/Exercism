// Package weather: A package to forecast the weather.
package weather

// CurrentCondition: Defines the current condition.
var CurrentCondition string

// CurrentLocation:  Defines the current location.
var CurrentLocation string

// Forecast: Gives the forcast for a given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
