# 15 — Pointers

**Code:** `15_pointers/pointers.go`

## Definition

A pointer stores the **memory address** of another variable.

| Syntax | Meaning |
|--------|---------|
| `*T` | type: pointer to T (`var p *int`) |
| `&x` | address of `x` |
| `*p` | value at that address (dereference) |
| `new(T)` | allocate a T, return `*T` (zero value) |

Zero value of a pointer is `nil`. Dereferencing `nil` panics.

## Why use

Go passes arguments **by value** (a copy). To change the caller’s variable, pass a pointer. Also used for large structs (avoid copying) and optional values (`*string` can be nil).

## Advantages

- Update the original from a function (`increment`, `swap`).
- No pointer arithmetic (safer than C).
- Garbage collected — no `free`.
- Struct fields via pointer use `.` (`userPtr.age`), not `->`.

## Disadvantages

- Extra mental load (`&` vs `*`).
- Nil pointer dereference crashes.
- Sharing one address from many goroutines needs a mutex.
- Escaping pointers can push data to the heap (slight cost).

## How to do it in Go

```go
num := 1
p := &num
fmt.Println(*p) // 1
*p = 20         // num is now 20

func changeByValue(n int) { n = 5 }      // copy — caller unchanged
func changeByPointer(n *int) { *n = 5 }  // same memory

b := 1
changeByPointer(&b) // b is 5
```

## In Python

Python has **no pointers**. Assignment binds names to objects.

- Immutable (`int`, `str`): `def f(n): n = 5` does not change the caller.
- Mutable (`list`, `dict`): `def f(xs): xs.append(1)` **does** change the caller.

```python
def increment(box):
    box[0] += 1          # fake "pointer" with a 1-item list

score = [10]
increment(score)
print(score[0])          # 11
```

## In other languages

**C / C++** — real pointers plus arithmetic (`p++`). C++ also has references (`int&`).

**Java / JavaScript / C#** — objects are references; primitives are copied. You cannot take `&int` in Java.

**Rust** — `&T` / `&mut T` borrows with a compiler-checked lifetime.

**TypeScript** — like JS; no address-of operator.
