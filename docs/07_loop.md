# 07 — Loops

**Code:** `07_loop/main.go`

## Definition

A loop repeats a block of code. **Go has only `for`**. You use it as a C-style loop, a while-loop, an infinite loop, or with `range` over a collection.

## Why use

Counting, walking lists, polling until a condition, skipping items (`continue`), stopping early (`break`). Almost every program needs repetition.

## Advantages

- One keyword to learn (`for`).
- `range` is safe (no off-by-one on slices/maps).
- `break` / `continue` work as expected.
- No `while` / `do-while` to remember.

## Disadvantages

- No `while` keyword — beginners look for it.
- Infinite `for { }` without `break` hangs the program.
- Classic `for i := 0; i < n; i++` is easy to get wrong on arrays.

## How to do it in Go

```go
for i := 1; i <= 5; i++ {        // classic
	fmt.Println(i)
}

for n > 0 {                       // while-style
	n--
}

for {                             // infinite
	if x > 3 {
		break
	}
}

for i := 1; i <= 5; i++ {
	if i%2 == 0 {
		continue
	}
}

for index, fruit := range fruits {
	fmt.Println(index, fruit)
}
```

## In Python

```python
for i in range(1, 6):
    print(i)

n = 3
while n > 0:
    print(n)
    n -= 1

for fruit in ["apple", "banana", "mango"]:
    print(fruit)

for i, fruit in enumerate(["apple", "banana"]):
    print(i, fruit)
```

Python has both `for` and `while`. `range(1, 6)` is 1…5.

## In other languages

**JavaScript**

```js
for (let i = 1; i <= 5; i++) {}
while (n > 0) { n--; }
for (const fruit of fruits) {}
fruits.forEach((fruit, i) => {});
```

**Java** — `for`, `while`, `do-while`, enhanced `for (String f : fruits)`.

Go 1.22+ also allows `for i := range 5` (0..4) — see lesson 11.
