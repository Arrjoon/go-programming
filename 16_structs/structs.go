package main

import (
	"fmt"
	"time"
)

// ========== WHAT IS A STRUCT? ==========
// A struct is a custom type that groups related fields (data) together.
// Think of a Python class, but data and methods are written separately.
//
// Python:
//   class Order:
//       def __init__(self, id, amount, status):
//           self.id = id
//           self.amount = amount
//           self.status = status
//       def change_status(self, status):
//           self.status = status
//
// Go:
//   type order struct { id, amount, status }
//   func newOrder(...) *order { ... }          ← constructor (no __init__)
//   func (o *order) changeStatus(...) { ... }  ← method (o is like self)
//
// Notes:
//  - Go has no class keyword. struct + methods = object-style code.
//  - Go has no real constructor. We write a function, often named newX / NewX.
//  - If you do not set a field, it gets the zero value:
//      int/float → 0, string → "", bool → false, pointer/time → zero/nil
//  - (o order)  value receiver  → works on a COPY (like pass by value)
//  - (o *order) pointer receiver → can change the original (like self)

// ========== STRUCT TYPE (the "class") ==========
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
}

// ========== CONSTRUCTOR (Python __init__) ==========
// Go does not have constructors. This function builds and returns an order.
func newOrder(id string, amount float32, status string) *order {
	myOrder := order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
	}
	return &myOrder
}

// ========== METHODS (Python class functions) ==========
// pointer receiver — can change the original order (use * when you need to update)
func (o *order) changeStatus(status string) {
	o.status = status
}

// value receiver — only reads; does not need to change the order
func (o order) getAmount() float32 {
	return o.amount
}

func (o order) info() string {
	return fmt.Sprintf("order %s: %.2f (%s)", o.id, o.amount, o.status)
}

// ========== ANOTHER STRUCT + EMBEDDING (like using one class inside another) ==========
type person struct {
	name string
	age  int
}

type employee struct {
	person // embedded — employee "has" name and age
	role   string
	salary int
}

func newPerson(name string, age int) person {
	return person{name: name, age: age}
}

func (p person) greet() {
	fmt.Println("hi, I am", p.name)
}

func (e employee) work() {
	fmt.Println(e.name, "works as", e.role) // name comes from embedded person
}

func main() {
	// ----- 1. struct literal (set fields by name) -----
	myOrder := order{
		id:     "1",
		amount: 50,
		status: "received",
	}
	myOrder.createdAt = time.Now()
	fmt.Println("literal:", myOrder)

	// ----- 2. another value -----
	myOrder2 := order{
		id:        "2",
		amount:    200,
		status:    "delivered",
		createdAt: time.Now(),
	}
	fmt.Println("second order:", myOrder2)

	// ----- 3. call methods (like obj.method() in Python) -----
	myOrder.changeStatus("paid")
	fmt.Println("amount:", myOrder.getAmount())
	fmt.Println("info:", myOrder.info())
	fmt.Println("after status change:", myOrder)

	// ----- 4. constructor function (like Order(...)) -----
	o3 := newOrder("3", 99.5, "pending")
	fmt.Println("from constructor:", o3)
	o3.changeStatus("shipped")
	fmt.Println("o3 info:", o3.info())

	// ----- 5. empty struct → all zero values -----
	var empty order
	fmt.Println("zero values:", empty)
	fmt.Println("empty id:", empty.id, "amount:", empty.amount, "status:", empty.status)

	// ----- 6. set fields one by one -----
	var o4 order
	o4.id = "4"
	o4.amount = 10
	o4.status = "new"
	fmt.Println("set one by one:", o4)

	// ----- 7. pointer to a struct -----
	p := &myOrder2
	fmt.Println("via pointer:", p.id, p.status)
	p.changeStatus("returned")
	fmt.Println("myOrder2 after pointer:", myOrder2.status)

	// ----- 8. person like a small Python class -----
	user := newPerson("arjun", 25)
	user.greet()
	fmt.Println("user:", user.name, user.age)

	// ----- 9. embedded struct (composition, not inheritance) -----
	emp := employee{
		person: person{name: "maya", age: 22},
		role:   "dev",
		salary: 50000,
	}
	emp.greet() // method from person still works
	emp.work()
	fmt.Println("employee name:", emp.name, "role:", emp.role)

	// ----- 10. anonymous struct (one-off, no type name) -----
	book := struct {
		title string
		year  int
	}{
		title: "The Go Programming Language",
		year:  2015,
	}
	fmt.Println("anonymous struct:", book)

	// ----- 11. slice of structs -----
	orders := []order{
		{id: "a", amount: 10, status: "new"},
		{id: "b", amount: 20, status: "paid"},
	}
	fmt.Println("slice of structs:")
	for _, item := range orders {
		fmt.Println(" ", item.info())
	}

	// ----- 12. compare structs (same fields + same types, no slices/maps inside) -----
	a := person{name: "ram", age: 20}
	b := person{name: "ram", age: 20}
	c := person{name: "sita", age: 20}
	fmt.Println("a == b?", a == b)
	fmt.Println("a == c?", a == c)
}
