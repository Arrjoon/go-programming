package main

import "fmt"

// functions: reusable blocks of code
// declare at package level, then call from main (or from other functions)

// ========== 1. NO PARAMS, NO RETURN ==========
func greet() {
	fmt.Println("hello from greet")
}

// ========== 2. PARAMETERS, NO RETURN ==========
func sayHello(name string) {
	fmt.Println("hello", name)
}

// ========== 3. SAME-TYPE PARAMS (write the type once) ==========
func add(a, b int) int {
	return a + b
}

// ========== 4. DIFFERENT-TYPE PARAMS + ONE RETURN ==========
func intro(name string, age int) string {
	return fmt.Sprintf("%s is %d years old", name, age)
}

// ========== 5. MULTIPLE RETURN VALUES ==========
// Go often returns (value, error). Here we return three strings.
func getLanguages() (string, string, string) {
	return "golang", "javascript", "c"
}

// ========== 6. NAMED RETURNS (bare return uses the names) ==========
func divide(a, b float64) (quot float64, ok bool) {
	if b == 0 {
		return 0, false
	}
	quot = a / b
	ok = true
	return
}

// ========== 7. VARIADIC FUNCTIONS ==========
// ...type means "zero or more arguments of this type"
// inside the function, the name is a slice (nums is []int)

// 7a. only variadic args
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// 7b. fixed params first, variadic LAST (only the last param can use ...)
func join(sep string, parts ...string) string {
	if len(parts) == 0 {
		return ""
	}
	result := parts[0]
	for _, p := range parts[1:] {
		result += sep + p
	}
	return result
}

// 7c. required first value + extra optional numbers
func maxOf(first int, rest ...int) int {
	m := first
	for _, n := range rest {
		if n > m {
			m = n
		}
	}
	return m
}

// 7d. mixed types with ...any (like fmt.Println)
func printAll(values ...any) {
	for i, v := range values {
		fmt.Println(i, v)
	}
}

// 7e. pass a variadic list into another variadic function
func printSum(label string, nums ...int) {
	fmt.Println(label, sum(nums...))
}

// 7f. variadic of functions — call every function you pass in
func runAll(fns ...func()) {
	for _, fn := range fns {
		fn()
	}
}

// ========== 8. FUNCTION TYPE (a name for a function signature) ==========
type mathFunc func(int, int) int

func mul(a, b int) int {
	return a * b
}

func sub(a, b int) int {
	return a - b
}

// apply takes another function and uses it
func apply(a, b int, fn mathFunc) int {
	return fn(a, b)
}

// ========== 9. PASS A FUNCTION AS AN ARGUMENT ==========
// processIt receives a function, then calls it
func processIt(fn func(a int) int) {
	result := fn(1)
	fmt.Println("processIt got:", result)
}

// applyToEach runs fn on every item in a slice
func applyToEach(nums []int, fn func(int) int) []int {
	out := make([]int, len(nums))
	for i, n := range nums {
		out[i] = fn(n)
	}
	return out
}

// filter keeps items where fn returns true
func filter(nums []int, fn func(int) bool) []int {
	var out []int
	for _, n := range nums {
		if fn(n) {
			out = append(out, n)
		}
	}
	return out
}

// ========== 10. RETURN A FUNCTION FROM A FUNCTION ==========
func makeAdder(x int) func(int) int {
	// the returned function "remembers" x (this is a closure)
	return func(y int) int {
		return x + y
	}
}

func makeGreeter(prefix string) func(string) {
	return func(name string) {
		fmt.Println(prefix, name)
	}
}

// ========== 11. CLOSURE — inner function uses outer variables ==========
func counter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

// ========== 12. RECURSIVE FUNCTION (calls itself) ==========
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// ========== 13. POINTER PARAMETER (can change the caller's value) ==========
func increment(n *int) {
	*n = *n + 1
}

// ========== 14. PASS SLICE / MAP (they already share the backing data) ==========
func appendHello(names []string) {
	names[0] = "hello " + names[0]
}

func setRole(user map[string]string, role string) {
	user["role"] = role
}

// ========== 15. DEFER — run this function just before the outer one returns ==========
func withDefer() {
	defer fmt.Println("defer: runs last")
	fmt.Println("withDefer: runs first")
}

// ========== 16. ANONYMOUS FUNCTION assigned to a package-level var ==========
var double = func(n int) int {
	return n * 2
}

func main() {
	// 1. no params
	greet()

	// 2. one param
	sayHello("arjun")

	// 3. same-type params + return
	fmt.Println("add:", add(12, 13))

	// 4. different params + return
	fmt.Println(intro("maya", 22))

	// 5. multiple returns — use all, or ignore some with _
	lang1, lang2, lang3 := getLanguages()
	fmt.Println("all langs:", lang1, lang2, lang3)
	first, second, _ := getLanguages()
	fmt.Println("ignore third:", first, second)
	fmt.Println(getLanguages())

	// 6. named returns
	q, ok := divide(10, 2)
	fmt.Println("10 / 2 =", q, "ok?", ok)
	_, ok = divide(10, 0)
	fmt.Println("10 / 0 ok?", ok)

	// 7. variadic functions
	// call with many values, with none, or unpack a slice with ...
	fmt.Println("sum 1,2,3:", sum(1, 2, 3))
	fmt.Println("sum none:", sum())
	more := []int{4, 5, 6}
	fmt.Println("sum slice:", sum(more...))

	fmt.Println("join:", join("-", "go", "is", "fun"))
	fmt.Println("join one:", join("-", "only"))
	fmt.Println("join none:", join("-"))
	words := []string{"a", "b", "c"}
	fmt.Println("join slice:", join(",", words...))

	fmt.Println("maxOf:", maxOf(3, 9, 2, 15, 7))
	fmt.Println("maxOf one:", maxOf(8))

	fmt.Println("printAll mixed types:")
	printAll("arjun", 25, true, 3.14)

	printSum("total", 10, 20, 30)
	printSum("from slice", more...)

	runAll(
		func() { fmt.Println("first task") },
		func() { fmt.Println("second task") },
		func() { fmt.Println("third task") },
	)

	// 8 + 9. pass one function into another
	fmt.Println("apply add:", apply(10, 3, add))
	fmt.Println("apply mul:", apply(10, 3, mul))
	fmt.Println("apply sub:", apply(10, 3, sub))

	// pass an anonymous function (no name)
	processIt(func(a int) int {
		return a * 10
	})

	// save a function in a variable, then pass it
	fn := func(a int) int {
		return 2
	}
	processIt(fn)

	// pass a named function that matches the signature
	processIt(double)

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("double each:", applyToEach(nums, double))
	fmt.Println("square each:", applyToEach(nums, func(n int) int {
		return n * n
	}))
	fmt.Println("only even:", filter(nums, func(n int) bool {
		return n%2 == 0
	}))

	// 10. function that returns a function
	add5 := makeAdder(5)
	fmt.Println("add5(10):", add5(10))
	add100 := makeAdder(100)
	fmt.Println("add100(10):", add100(10))

	greetDev := makeGreeter("hi")
	greetDev("golang")

	// 11. closure keeps its own state
	next := counter()
	fmt.Println("counter:", next(), next(), next())
	other := counter()
	fmt.Println("new counter:", other())

	// 12. recursion
	fmt.Println("factorial 5:", factorial(5))

	// 13. pointer param
	x := 10
	increment(&x)
	fmt.Println("after increment:", x)

	// 14. slice / map as args
	names := []string{"arjun", "maya"}
	appendHello(names)
	fmt.Println("slice after func:", names)

	user := map[string]string{"name": "arjun"}
	setRole(user, "dev")
	fmt.Println("map after func:", user)

	// 15. defer
	withDefer()

	// 16. immediately-invoked anonymous function
	func(msg string) {
		fmt.Println("iife:", msg)
	}("runs right away")

	// 17. function value you can replace
	var op mathFunc
	op = add
	fmt.Println("op is add:", op(2, 3))
	op = mul
	fmt.Println("op is mul:", op(2, 3))
}
