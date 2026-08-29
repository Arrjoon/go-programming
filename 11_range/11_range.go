package main

import "fmt"

// range loops over a slice, array, map, or string
func main() {
	// ========== SLICE ==========
	nums := []int{10, 20, 30}

	// 1. index + value
	fmt.Println("slice index + value:")
	for i, v := range nums {
		fmt.Println(i, v)
	}

	// 2. index only
	fmt.Println("slice index only:")
	for i := range nums {
		fmt.Println(i, "->", nums[i])
	}

	// 3. value only (ignore index with _)
	fmt.Println("slice value only:")
	for _, v := range nums {
		fmt.Println(v)
	}

	// 4. ignore both (run once per item)
	fmt.Println("slice ignore both:")
	count := 0
	for range nums {
		count++
	}
	fmt.Println("how many items:", count)

	// 5. empty / nil slice — loop body never runs
	var empty []string
	for i, v := range empty {
		fmt.Println("this will not print", i, v)
	}
	fmt.Println("empty slice done")

	// ========== ARRAY ==========
	fruits := [3]string{"apple", "banana", "mango"}

	// 6. index + value
	fmt.Println("array index + value:")
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}

	// 7. index only
	fmt.Println("array index only:")
	for i := range fruits {
		fmt.Println(i)
	}

	// 8. value only
	fmt.Println("array value only:")
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	// 9. 2D array (nested range)
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2D array:")
	for row, cols := range matrix {
		for col, value := range cols {
			fmt.Println("row", row, "col", col, "=", value)
		}
	}

	// ========== MAP ==========
	ages := map[string]int{
		"alice": 24,
		"bob":   30,
		"carol": 22,
	}

	// 10. key + value
	fmt.Println("map key + value:")
	for name, age := range ages {
		fmt.Println(name, "->", age)
	}

	// 11. keys only
	fmt.Println("map keys only:")
	for name := range ages {
		fmt.Println(name)
	}

	// 12. values only
	fmt.Println("map values only:")
	for _, age := range ages {
		fmt.Println(age)
	}

	// 13. ignore both
	fmt.Println("map ignore both:")
	n := 0
	for range ages {
		n++
	}
	fmt.Println("how many keys:", n)

	// 14. map of slices
	groups := map[string][]string{
		"frontend": {"html", "css"},
		"backend":  {"go", "sql"},
	}
	fmt.Println("map of slices:")
	for group, items := range groups {
		for i, item := range items {
			fmt.Println(group, i, item)
		}
	}

	// ========== STRING ==========
	// range on a string gives byte index + rune (not a byte)
	word := "Go"
	fmt.Println("string index + rune:")
	for i, r := range word {
		fmt.Println("index", i, "rune", r, "char", string(r))
	}

	// 15. Unicode — index jumps by bytes, value is the full character
	name := "नेपाल"
	fmt.Println("unicode string:")
	for i, r := range name {
		fmt.Println("byte index", i, "char", string(r))
	}

	// 16. index only (byte indexes of each character start)
	fmt.Println("string index only:")
	for i := range name {
		fmt.Println(i)
	}

	// 17. rune only
	fmt.Println("string rune only:")
	for _, r := range name {
		fmt.Println(string(r))
	}

	// 18. bytes vs range — ranging a []byte is per byte
	fmt.Println("range over []byte (each byte):")
	for i, b := range []byte("ने") {
		fmt.Println(i, b)
	}

	// 19. range over []rune (each character)
	fmt.Println("range over []rune (each char):")
	for i, r := range []rune("नेपाल") {
		fmt.Println(i, string(r))
	}

	// ========== EXTRA ==========
	// 20. range over an integer (Go 1.22+) — 0, 1, 2, ... n-1
	fmt.Println("range over int:")
	for i := range 5 {
		fmt.Println(i)
	}

	// 21. break and continue
	fmt.Println("break / continue:")
	for i, v := range nums {
		if v == 20 {
			continue
		}
		if v == 30 {
			break
		}
		fmt.Println(i, v)
	}
}
