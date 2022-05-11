// Package weather has a city and condition variable, and a function formatting them into a string weather forecast.
package weather
// CurrentCondition stores the current weather as a string.
var CurrentCondition string
// CurrentLocation stores the current location as a string.
var CurrentLocation string
// Forecast accepts a city and condition, returning a string with the entered forecast formatted.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
