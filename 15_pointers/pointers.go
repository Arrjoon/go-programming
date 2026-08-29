package main

import "fmt"

// ========== WHAT IS A POINTER? ==========
// A pointer is a variable that stores the MEMORY ADDRESS of another variable.
//
//   num := 10
//   &num  → address of num   (where it lives in memory)
//   *p    → value at that address (the real number)
//
// Syntax:
//   *T     type: "pointer to T"     var p *int
//   &x     address of x
//   *p     dereference: read/write the value p points to
//   new(T) create a T in memory and return a pointer to it
//
// Why use pointers?
//   Pass by value  → function gets a COPY. Changing it does not change main.
//   Pass by pointer (reference) → function gets the ADDRESS. Changing *p
//   changes the original variable.

// ========== PASS BY VALUE (copy) ==========
// num is a copy. setting it to 5 does NOT change the caller's variable.
func changeByValue(num int) {
	num = 5
	fmt.Println("inside changeByValue, copy is", num)
}

// ========== PASS BY POINTER (same memory) ==========
// num is *int — an address. *num = 5 writes into the original variable.
func changeByPointer(num *int) {
	*num = 5
	fmt.Println("inside changeByPointer, value is", *num)
	fmt.Println("inside changeByPointer, address is", num)
}

// swap two ints using pointers
func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

// increment through a pointer
func increment(n *int) {
	*n = *n + 1
}

func main() {
	// ----- 1. normal variable + address -----
	num := 1
	fmt.Println("value of num:", num)
	fmt.Println("address of num (&num):", &num)

	// ----- 2. declare a pointer (zero value is nil) -----
	var p *int
	fmt.Println("empty pointer:", p)
	fmt.Println("is nil?", p == nil)

	// ----- 3. point p at num -----
	p = &num
	fmt.Println("p holds address:", p)
	fmt.Println("value at p (*p):", *p)

	// ----- 4. change num THROUGH the pointer -----
	*p = 20
	fmt.Println("after *p = 20, num is", num)

	// ----- 5. short declare a pointer -----
	x := 10
	ptr := &x
	fmt.Println("ptr:", ptr, "value:", *ptr)

	// ----- 6. new(T) — Go creates the value and gives you a pointer -----
	n := new(int) // *n starts at 0
	fmt.Println("new int:", *n)
	*n = 42
	fmt.Println("after set:", *n)

	// ----- 7. pointer to a string -----
	name := "golang"
	namePtr := &name
	fmt.Println("string value:", *namePtr)
	*namePtr = "go"
	fmt.Println("name after pointer write:", name)

	// ----- 8. pointer to a struct (use . directly, no extra *) -----
	type person struct {
		name string
		age  int
	}
	user := person{name: "arjun", age: 25}
	userPtr := &user
	fmt.Println("struct via pointer:", userPtr.name, userPtr.age)
	userPtr.age = 26
	fmt.Println("user after pointer write:", user)

	// ----- 9. PASS BY VALUE vs PASS BY POINTER -----
	a := 1
	fmt.Println("--- pass by value ---")
	fmt.Println("before:", a)
	changeByValue(a) // sends a COPY
	fmt.Println("after (still 1):", a)

	b := 1
	fmt.Println("--- pass by pointer ---")
	fmt.Println("before:", b)
	fmt.Println("address we pass:", &b)
	changeByPointer(&b) // sends the ADDRESS
	fmt.Println("after (now 5):", b)

	// ----- 10. swap and increment -----
	left, right := 7, 3
	fmt.Println("before swap:", left, right)
	swap(&left, &right)
	fmt.Println("after swap:", left, right)

	score := 10
	increment(&score)
	increment(&score)
	fmt.Println("score after increment:", score)
}
