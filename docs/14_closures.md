# 14 — Closures

**Code:** `14_closures/closures.go`

## Definition

A closure is a function that **remembers variables** from the function around it. Those variables stay alive after the outer function returns. The inner function “closes over” that data.

Properties: it can read and write outer vars; each call to the outer function gets its **own** copy of the state; the closure is a value you can store, pass, or return.

## Why use

Counters, bank-style balances, filters (`greaterThan(5)`), callbacks, configuration, and avoiding globals. The lesson file shows `counter`, `addN`, `bankAccount`, and `runningTotal`.

## Advantages

- Private state without a struct (sometimes).
- Factory functions: `add5 := addN(5)`.
- Two functions can share one captured variable (deposit + balance).
- Works as a callback (`applyTwice`, `keep`).

## Disadvantages

- Hidden state is harder to debug than a struct field.
- Loop + goroutine capture used to share one `i` (Go &lt; 1.22). Pass `i` as an argument to be safe.
- Closures that capture large objects keep them alive (memory).
- Not a replacement for clear types when the state is complex.

## How to do it in Go

```go
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

increment := counter()
increment() // 1
increment() // 2

other := counter() // independent count
```

## In Python

```python
def counter():
    count = 0
    def inner():
        nonlocal count
        count += 1
        return count
    return inner

increment = counter()
print(increment())  # 1
```

Python needs `nonlocal` to assign to an outer variable. `lambda` can close over reads.

## In other languages

**JavaScript**

```js
function counter() {
  let count = 0;
  return () => ++count;
}
```

Closures are the backbone of JS callbacks and hooks.

**Java** — lambdas can capture `final` or effectively-final locals; for mutable state use an array or a class.

**Rust** — `move` closures; borrow checker rules apply.
