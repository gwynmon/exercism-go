// Package weather provides weather forecast.
package weather

var (
    // CurrentCondition is a variable representing current condition.
	CurrentCondition string
    // CurrentLocation is a variable representing current location.
	CurrentLocation  string
)

// Forecast takes city location, current condition and assign it to variables.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
