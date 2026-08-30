# 17 — Struct embedding

**Code:** `17_struct_embedding/embedding.go`

## Definition

Embedding puts one struct **inside another without a field name**. Fields and methods of the inner type are **promoted** — you can write `emp.name` and `emp.greet()` as if they belonged to `employee`.

This is **composition** (“has a”), not Python-style inheritance (“is a”).

```go
type employee struct {
	person   // embedded
	role string
}
```

A **named** field (`user person`) does **not** promote — you must use `s.user.name`.

## Why use

Reuse fields and methods without copying them. Nested types (manager embeds employee embeds person), contacts with address, optional pointer embedding (`*person`).

## Advantages

- Promoted fields and methods — less boilerplate.
- Full path still works: `e.person.name`.
- Outer method **wins** if names clash; inner stays at `e.person.info()`.
- Embed more than one struct (`person` + `address`).
- Pointer embed for optional / shared inner values.

## Disadvantages

- Not real inheritance (no `super`, no polymorphism by itself).
- Name clashes can surprise you (outer hides inner).
- Embedding many types makes the API unclear (“where does `name` live?”).
- Promoted methods do not automatically satisfy an interface in a “subclass” sense — the outer type’s method set is what matters.

## How to do it in Go

```go
type person struct{ name string; age int }
func (p person) greet() { fmt.Println("hi", p.name) }

type employee struct {
	person
	role   string
	salary int
}

e := employee{person: person{name: "maya", age: 22}, role: "dev"}
e.greet()          // promoted
e.name             // promoted
e.person.info()    // inner, if names clash
```

## In Python

```python
class Person:
    def __init__(self, name, age):
        self.name = name
        self.age = age
    def greet(self):
        print("hi", self.name)

class Employee(Person):          # inheritance
    def __init__(self, name, age, role):
        super().__init__(name, age)
        self.role = role

# composition (closer to Go):
class Staff:
    def __init__(self, user: Person, role: str):
        self.user = user
        self.role = role
```

Go embedding ≈ composition + automatic forwarding of names.

## In other languages

**Java / C# / TypeScript** — `extends` / `implements` for inheritance; composition is a field (`private Person person`). No promotion unless you write wrappers.

**Rust** — no embedding; you compose fields and deref or write methods.

**Kotlin** — `class Employee(p: Person) : Person by p` (delegation) is similar in spirit.
