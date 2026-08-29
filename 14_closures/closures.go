package main

import "fmt"

// ========== WHAT IS A CLOSURE? ==========
// A closure is a function that uses variables from the function around it.
// Those outer variables stay alive after the outer function returns.
// The inner function "closes over" (remembers) that data.
//
// Properties:
//  1. It can READ outer variables.
//  2. It can WRITE / change outer variables.
//  3. Those variables live as long as the closure lives.
//  4. Each call to the outer function gets its OWN copy of the data.
//  5. You can assign a closure, pass it, or return it (it is a value).

// ========== 1. BASIC — return a function that remembers count ==========
func counter() func() int {
	var count int = 0

	// this inner function is the closure
	return func() int {
		count += 1
		return count
	}
}

// ========== 2. CLOSURE THAT TAKES AN ARGUMENT ==========
// addN remembers n, then adds it to whatever you pass later
func addN(n int) func(int) int {
	return func(x int) int {
		return x + n
	}
}

// ========== 3. CLOSURE THAT RETURNS TWO FUNCTIONS (share the same data) ==========
func bankAccount(start int) (deposit func(int), balance func() int) {
	money := start

	deposit = func(amount int) {
		money += amount
	}
	balance = func() int {
		return money
	}
	return
}

// ========== 4. CLOSURE USED AS A CALLBACK (passed to another function) ==========
func applyTwice(n int, fn func(int) int) int {
	return fn(fn(n))
}

// ========== 5. FILTER WITH A CLOSURE (the test function remembers a limit) ==========
func greaterThan(limit int) func(int) bool {
	return func(n int) bool {
		return n > limit
	}
}

func keep(nums []int, test func(int) bool) []int {
	var out []int
	for _, n := range nums {
		if test(n) {
			out = append(out, n)
		}
	}
	return out
}

// ========== 6. ADDER + SUBTRACTOR that share one running total ==========
func runningTotal() (add func(int), sub func(int), get func() int) {
	total := 0
	add = func(n int) { total += n }
	sub = func(n int) { total -= n }
	get = func() int { return total }
	return
}

func main() {
	// ----- definition in action -----
	// increment "closes over" count inside counter()
	increment := counter()

	fmt.Println("counter:", increment()) // 1
	fmt.Println("counter:", increment()) // 2
	fmt.Println("counter:", increment()) // 3  ← same count, still remembered

	// property 4: a new call to counter() gets a NEW count (starts at 0 again)
	other := counter()
	fmt.Println("other counter:", other())     // 1
	fmt.Println("first counter:", increment()) // 4  ← first one is unchanged

	// ----- addN remembers n -----
	add5 := addN(5)
	add10 := addN(10)
	fmt.Println("7 + 5 =", add5(7))
	fmt.Println("7 + 10 =", add10(7))

	// ----- two functions share the same captured variable -----
	deposit, balance := bankAccount(100)
	fmt.Println("start:", balance())
	deposit(50)
	deposit(25)
	fmt.Println("after deposits:", balance())

	// ----- pass a closure into another function -----
	double := func(n int) int {
		return n * 2
	}
	fmt.Println("applyTwice 3:", applyTwice(3, double)) // (3*2)*2 = 12

	// anonymous closure passed directly
	fmt.Println("applyTwice +1:", applyTwice(10, func(n int) int {
		return n + 1
	})) // 12

	// ----- closure as a filter -----
	nums := []int{1, 4, 8, 12, 3}
	fmt.Println("greater than 5:", keep(nums, greaterThan(5)))
	fmt.Println("greater than 10:", keep(nums, greaterThan(10)))

	// ----- local closure (declared inside main) -----
	message := "hi"
	say := func(name string) {
		// reads the outer variable message
		fmt.Println(message, name)
	}
	say("arjun")
	message = "hello" // change outer var — closure sees the new value
	say("maya")

	// ----- write to an outer variable from a local closure -----
	score := 0
	addPoint := func() {
		score++
	}
	addPoint()
	addPoint()
	fmt.Println("score:", score)

	// ----- shared running total -----
	add, sub, get := runningTotal()
	add(10)
	add(5)
	sub(3)
	fmt.Println("running total:", get()) // 12

	// ----- immediately-invoked closure -----
	func(prefix string) {
		fmt.Println(prefix, "this closure ran at once")
	}("note:")

	// ----- loop + closure: each function remembers its own i -----
	var printers []func()
	for i := 1; i <= 3; i++ {
		n := i // own copy for this loop step
		printers = append(printers, func() {
			fmt.Println("loop value:", n)
		})
	}
	for _, p := range printers {
		p()
	}
}

// Closure remembers variables from its outer function even after the outer function finishes.
// It allows a function to maintain state between multiple calls.
// It can read and modify the captured outer variables.
// Each call to the outer function creates separate, independent state.
// It helps avoid global variables by keeping data private to the closure.
// A closure can be returned, stored in a variable, or passed as an argument.
// It is useful for counters, callbacks, filters, configuration, and stateful functions.
