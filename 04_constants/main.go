package main

import "fmt"

// Constants: values that never change.
func main() {
	const pi = 3.14159
	const greeting = "Hello"

	fmt.Println(greeting, "pi is", pi)

	// Typed constant
	const daysInWeek int = 7
	fmt.Println("Days in a week:", daysInWeek)

	// Multiple constants
	const (
		statusOK    = 200
		statusError = 500
	)
	fmt.Println("OK:", statusOK, "Error:", statusError)

	// iota: auto-incrementing constants (0, 1, 2, ...)
	const (
		Sunday = iota
		Monday
		Tuesday
	)
	fmt.Println("Sunday:", Sunday, "Monday:", Monday, "Tuesday:", Tuesday)
}
