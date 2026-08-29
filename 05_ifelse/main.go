package main

import "fmt"

// if / else: run code based on conditions.
func main() {
	age := 18

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	score := 75

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: D or below")
	}

	// if with a short statement (variable only lives inside if/else)
	if n := 10; n%2 == 0 {
		fmt.Println(n, "is even")
	} else {
		fmt.Println(n, "is odd")
	}
}
