package main

import "fmt"

// switch: choose one branch from many options.
func main() {
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Weekend")
	}

	// Multiple values in one case
	letter := "a"
	switch letter {/
	case "a", "e", "i", "o", "u":
		fmt.Println(letter, "is a vowel")
	default:
		fmt.Println(letter, "is a consonant")
	}

	// switch without a value (like if / else if)
	hour := 14
	switch {
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 18:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
