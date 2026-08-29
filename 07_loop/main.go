package main

import "fmt"

// Loops: Go has only for (used in several ways).
func main() {
	// Classic for loop
	fmt.Println("Count 1 to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// While-style loop (condition only)
	fmt.Println("Countdown:")
	n := 3
	for n > 0 {
		fmt.Println(n)
		n--
	}

	// Infinite loop with break
	fmt.Println("Break example:")
	x := 0
	for {
		x++
		if x > 3 {
			break
		}
		fmt.Println(x)
	}

	// continue: skip this round
	fmt.Println("Odd numbers 1 to 5:")
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// Range over a slice (array-like list)
	fmt.Println("Fruits:")
	fruits := []string{"apple", "banana", "mango"}
	for index, fruit := range fruits {
		fmt.Println(index, fruit)
	}
}
