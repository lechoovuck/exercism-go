// Package weather contains a function to show forecast and weather conditions.
package weather

var (
	// CurrentCondition describes the current weather conditions.
	CurrentCondition string
	// CurrentLocation stores the location with the provided condition.
	CurrentLocation string
)

// Forecast function shows the weather conditions for the provided location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
