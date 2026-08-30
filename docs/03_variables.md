# 03 — Variables

**Code:** `03_variables/main.go`

## Definition

A variable is a named place that holds a value. In Go you declare the name and (usually) the type. If you do not set a value, you get the **zero value** (`0`, `""`, `false`, `nil`).

## Why use

You need names for data that changes: user input, counters, results. Zero values mean you rarely need `null` checks for ints and strings.

## Advantages

- Compile-time type checks catch mistakes early.
- Zero values are safe defaults (no “uninitialized garbage” like C).
- `:=` is short inside functions.
- Unused variables are a **compile error** — less dead code.

## Disadvantages

- More typing than Python (`var name string = "Arjun"`).
- Unused imports and variables fail the build.
- `:=` vs `=` is easy to mix up (`=` is assign, `:=` is declare + assign).

## How to do it in Go

```go
var name string = "Arjun"   // explicit type
var age = 25                // type inferred
city := "Kathmandu"         // short declare (inside functions only)

var a, b int = 1, 2

var emptyInt int       // 0
var emptyString string // ""
var emptyBool bool     // false
```

`var` works at package level. `:=` works only inside a function.

## In Python

```python
name = "Arjun"          # no type required
age = 25
city = "Kathmandu"
a, b = 1, 2

empty_int = 0           # you choose the default yourself
empty_string = ""
empty_bool = False
```

Optional types (3.5+): `name: str = "Arjun"`.

## In other languages

**JavaScript**

```js
let name = "Arjun";     // can change
const age = 25;         // cannot reassign
var city = "Kathmandu"; // older, function-scoped — avoid
```

**Java**

```java
String name = "Arjun";
int age = 25;
int emptyInt = 0;       // locals must be set before use
```

**TypeScript** — `let city: string = "Kathmandu"`.
