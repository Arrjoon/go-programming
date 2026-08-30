# 05 — If / else

**Code:** `05_ifelse/main.go`

## Definition

`if` / `else if` / `else` runs different code depending on a boolean condition. Go has **no parentheses** around the condition and **requires braces**. You can declare a short variable that exists only inside the if/else.

## Why use

Branching is how programs make decisions: age checks, grades, even/odd, error handling (`if err != nil`).

## Advantages

- No parentheses — less noise than C/Java.
- Braces are required — fewer “dangling else” bugs.
- Short statement (`if n := 10; n%2 == 0`) keeps the variable local.
- Conditions must be `bool` (not `if 1` like C).

## Disadvantages

- No ternary `a ? b : c` (use a full `if` or a small helper).
- Easy to forget the extra `;` in the short-statement form.

## How to do it in Go

```go
if age >= 18 {
	fmt.Println("You are an adult.")
} else {
	fmt.Println("You are a minor.")
}

if score >= 90 {
	fmt.Println("Grade: A")
} else if score >= 80 {
	fmt.Println("Grade: B")
} else {
	fmt.Println("Grade: D or below")
}

if n := 10; n%2 == 0 {
	fmt.Println(n, "is even")
}
```

## In Python

```python
if age >= 18:
    print("You are an adult.")
else:
    print("You are a minor.")

if score >= 90:
    print("Grade: A")
elif score >= 80:
    print("Grade: B")
else:
    print("Grade: D or below")
```

Python uses `elif` and indentation instead of braces. Ternary exists: `"even" if n % 2 == 0 else "odd"`.

## In other languages

**JavaScript / TypeScript**

```js
if (age >= 18) {
  console.log("adult");
} else if (age >= 13) {
  console.log("teen");
} else {
  console.log("child");
}
const label = age >= 18 ? "adult" : "minor";
```

**Java** — same `if (cond) { }` as JS; condition must be boolean.
