# 13 — Variadic functions

**Code:** `13_variadic_functions/variadic.go`

## Definition

A variadic function accepts **zero or more** arguments of one type. The last parameter uses `...T`. Inside the function that name is a **slice** (`[]T`). `fmt.Println` is variadic (`...any`).

## Why use

When the caller should pass a list without building a slice first: `sum(1, 2, 3)`, `join("-", "a", "b")`, logging. You can still pass an existing slice with `nums...`.

## Advantages

- Call site stays short: `sum(3, 4, 5, 6)`.
- Zero arguments are allowed (`sum()` → 0).
- Unpack a slice: `sum(more...)`.
- Last-arg-only rule keeps the signature readable.

## Disadvantages

- Only the **last** parameter can be variadic.
- One variadic type (or `...any` if you mix types — you lose type safety).
- Easy to forget `...` when forwarding a slice (passes the slice as one item).
- A nil/empty `...` is a nil slice — handle `len == 0`.

## How to do it in Go

```go
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

sum(3, 4, 5, 6)
sum()
more := []int{1, 2, 3}
sum(more...)                 // unpack

func join(sep string, parts ...string) string { /* ... */ }
```

## In Python

```python
def sum_all(*nums):
    total = 0
    for n in nums:
        total += n
    return total

sum_all(3, 4, 5, 6)
sum_all(*[1, 2, 3])          # unpack

def join(sep, *parts):
    return sep.join(parts)
```

`*args` is the same idea. `**kwargs` has no direct Go equivalent.

## In other languages

**JavaScript**

```js
function sum(...nums) {
  return nums.reduce((a, b) => a + b, 0);
}
sum(3, 4, 5);
sum(...[1, 2, 3]);
```

**Java** — `int sum(int... nums)` (array inside).

**C** — `va_list` (unsafe, untyped). Go’s `...T` is typed.
