package main

import "fmt"

// Variables: declare and use named values.
func main() {
	// Declare with var and type
	var name string = "Arjun"
	fmt.Println("name:", name)

	// Type inferred by Go
	var age = 25
	fmt.Println("age:", age)

	// Short declaration (most common inside functions)
	city := "Kathmandu"
	fmt.Println("city:", city)

	// Multiple variables at once
	var a, b int = 1, 2
	fmt.Println("a + b =", a+b)

	// Zero values (default if not set)
	var emptyInt int
	var emptyString string
	var emptyBool bool
	fmt.Println("zero int:", emptyInt)
	fmt.Println("zero string:", emptyString)
	fmt.Println("zero bool:", emptyBool)
}
