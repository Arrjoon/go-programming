# 19 — Generics

**Code:** `19_generics/generics.go`

## Definition

Generics (Go 1.18+) let you write **one function or type** that works for many types, without `any` and without copy-paste.

```go
func printSlice[T any](items []T)
```

- `T` — type parameter (placeholder)
- `any` — constraint (which types T may be)
- Other constraints: `comparable`, unions (`~int | ~float64`)

The compiler often **infers** `T` from arguments (`printSlice(nums)`).

## Why use

Same algorithm for `[]int` and `[]string`: print, first/last, contains, sum, stack, queue, set, `result[T]`, a generic store. Keeps type safety — the result is `T`, not `any`.

## Advantages

- One implementation, many types.
- No type assertions at the call site.
- Constraints document what T can do (`==`, `+`).
- Generic structs (`stack[T]`) and interfaces (`store[T]`).
- `~int` accepts `type Age int`.

## Disadvantages

- Error messages can be harder to read.
- Not every problem needs generics (a plain interface is often clearer).
- Constraints cannot express every operation (no “has field X”).
- Methods cannot add extra type parameters beyond the receiver’s.
- Older Go (&lt; 1.18) cannot compile this code.

## How to do it in Go

```go
func first[T any](items []T) (T, bool) { /* ... */ }
func contains[T comparable](items []T, target T) bool { /* == */ }

type Number interface {
	~int | ~int64 | ~float64
}
func sum[T Number](nums []T) T { /* += */ }

type stack[T any] struct{ items []T }
func (s *stack[T]) push(v T) { s.items = append(s.items, v) }
```

## In Python

```python
from typing import TypeVar

T = TypeVar("T")

def first(items: list[T]) -> T:
    return items[0]

# runtime is still dynamic — hints are for checkers
def print_slice(items: list[T]) -> None:
    for x in items:
        print(x)
```

Python 3.12: `def first[T](items: list[T]) -> T`.

## In other languages

**TypeScript**

```ts
function printSlice<T>(items: T[]): void {}
class Stack<T> { push(v: T) {} }
```

**Java** — `List<T>`, `static <T> T first(List<T> items)` (type erasure at runtime).

**C++** — templates (`template<typename T>`), compiled per type.

**Rust** — `fn first<T>(items: &[T]) -> &T` with trait bounds (`T: PartialEq`).
