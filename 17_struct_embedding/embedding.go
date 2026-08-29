package main

import "fmt"

// ========== WHAT IS STRUCT EMBEDDING? ==========
// Embedding puts one struct INSIDE another without giving it a field name.
//
//   type employee struct {
//       person   // embedded — no field name
//       role string
//   }
//
// This is composition ("has a"), not Python-style inheritance ("is a").
//
// Python (inheritance):
//   class Person:
//       def greet(self): ...
//   class Employee(Person):
//       def work(self): ...
//
// Go (embedding):
//   type person struct { name string }
//   func (p person) greet() { ... }
//   type employee struct { person; role string }
//   // employee can use greet() and name as if they were its own
//
// Properties:
//  1. Fields of the inner struct are PROMOTED (emp.name works).
//  2. Methods of the inner struct are PROMOTED (emp.greet() works).
//  3. You can still use the full path: emp.person.name
//  4. If the outer type has the same method/field, the OUTER one wins.
//  5. You can embed more than one struct.

type person struct {
	name string
	age  int
}

func (p person) greet() {
	fmt.Println("hi, I am", p.name)
}

func (p person) info() {
	fmt.Println(p.name, "is", p.age)
}

// named field (NOT embedding) — must use emp.user.name
type staff struct {
	user person
	role string
}

// embedding — no field name, so emp.name and emp.greet() work
type employee struct {
	person
	role   string
	salary int
}

func newEmployee(name string, age int, role string, salary int) employee {
	return employee{
		person: person{name: name, age: age},
		role:   role,
		salary: salary,
	}
}

func (e employee) work() {
	fmt.Println(e.name, "works as", e.role)
}

// outer method with the same name hides the inner one
func (e employee) info() {
	fmt.Println(e.name, "is a", e.role, "age", e.age)
}

// embed more than one struct
type address struct {
	city    string
	country string
}

type contact struct {
	person
	address
	phone string
}

// nested embedding: manager embeds employee, which embeds person
type manager struct {
	employee
	teamSize int
}

// pointer embedding — useful when the inner value is optional / shared
type account struct {
	*person
	email string
}

func main() {
	// ----- 1. named field vs embedding -----
	s := staff{
		user: person{name: "ram", age: 20},
		role: "intern",
	}
	fmt.Println("named field: must use s.user.name ->", s.user.name)
	// fmt.Println(s.name)  // would not compile — name is not promoted

	e := newEmployee("maya", 22, "dev", 50000)
	fmt.Println("embedded: e.name ->", e.name)
	fmt.Println("full path: e.person.name ->", e.person.name)
	fmt.Println("own field: e.role ->", e.role)

	// ----- 2. promoted methods -----
	e.greet() // person.greet is promoted
	e.work()  // employee's own method

	// ----- 3. outer method wins if names clash -----
	e.info()         // employee.info, not person.info
	e.person.info()  // still can call the inner one

	// ----- 4. change promoted fields -----
	e.age = 23
	e.salary = 55000
	fmt.Println("after update:", e.name, e.age, e.salary)

	// ----- 5. embed two structs -----
	c := contact{
		person:  person{name: "arjun", age: 25},
		address: address{city: "Kathmandu", country: "Nepal"},
		phone:   "9800000000",
	}
	fmt.Println("contact name:", c.name)
	fmt.Println("contact city:", c.city)
	fmt.Println("contact phone:", c.phone)
	c.greet()

	// ----- 6. nested embedding -----
	m := manager{
		employee: newEmployee("sita", 30, "lead", 80000),
		teamSize: 5,
	}
	fmt.Println("manager name (from person):", m.name)
	fmt.Println("manager role (from employee):", m.role)
	fmt.Println("team size:", m.teamSize)
	m.greet()
	m.work()
	m.info()

	// ----- 7. pointer embedding -----
	p := &person{name: "hari", age: 28}
	acc := account{person: p, email: "hari@mail.com"}
	fmt.Println("account name:", acc.name)
	acc.greet()
	p.name = "hari bahadur"
	fmt.Println("account sees the change:", acc.name)

	// ----- 8. slice of embedded structs -----
	team := []employee{
		newEmployee("a", 21, "qa", 40000),
		newEmployee("b", 24, "dev", 50000),
	}
	fmt.Println("team:")
	for _, member := range team {
		fmt.Println(" ", member.name, member.role)
	}
}
