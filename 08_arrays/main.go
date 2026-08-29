package main

import "fmt"

// Arrays: fixed-size list of values of the same type.
func main() {
	// Declare an array of 5 ints (all start at 0)
	var numbers [5]int
	fmt.Println("empty array:", numbers)

	// Set values by index (starts at 0)
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	fmt.Println("after set:", numbers)
	fmt.Println("first item:", numbers[0])
	fmt.Println("length:", len(numbers))

	// Declare and fill in one line
	fruits := [3]string{"apple", "banana", "mango"}
	fmt.Println("fruits:", fruits)

	// Let Go count the size with [...]
	scores := [...]int{90, 85, 78, 92,23}
	fmt.Println("scores:", scores)
	fmt.Println("how many scores:", len(scores))

	// Loop through an array with for
	fmt.Println("print each fruit:")
	for i := 0; i < len(fruits); i++ {
		fmt.Println(i, fruits[i])
	}

	// Loop with range
	fmt.Println("print each score:")
	for index, value := range scores {
		fmt.Println(index, "->", value)
	}

	// 2D array (array of arrays)
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2D array:", matrix)
	fmt.Println("row 0, col 1:", matrix[0][1])
}
