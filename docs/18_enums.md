# 18 — Enums

**Code:** `18_enums/enums.go`

## Definition

Go has **no `enum` keyword**. You fake enumerated types with:

1. a **custom type** (`type OrderStatus int`)
2. a **const block** (often with `iota`, or with string literals)

That stops typos like `"delieverd"` from compiling when the function wants `OrderStatus`.

## Why use

Fixed sets of values: order status, roles, payment methods, log levels, environments, weekdays, permissions (bit flags). Add `String()`, `IsValid()`, and `ParseX()` when values come from users or APIs.

## Advantages

- Compiler checks the type (not a random `int`/`string` in theory — still validate input).
- `iota` numbers states automatically.
- String enums read well in logs and JSON.
- Bit flags (`PermRead | PermWrite`) pack many switches.
- You can attach methods (`CanEdit()`, `IsWeekend()`).

## Disadvantages

- Still just an `int` or `string` at runtime — `OrderStatus(99)` compiles.
- No exhaustiveness check on `switch` (always write `default`).
- Must write `String()` / parse yourself (or generate).
- `iota` in the same block as a manual `0` is easy to get wrong (bit flags).

## How to do it in Go

```go
type OrderStatus int
const (
	Received OrderStatus = iota
	Confirmed
	Delivered
)

type PayMethod string
const (
	PayEsewa  PayMethod = "esewa"
	PayKhalti PayMethod = "khalti"
)

func (s OrderStatus) String() string { /* switch */ }
```

See the lesson for roles, priority, ticket state machines, and `Permission` flags.

## In Python

```python
from enum import Enum, IntEnum, auto

class OrderStatus(IntEnum):
    RECEIVED = auto()
    CONFIRMED = auto()
    DELIVERED = auto()

class PayMethod(str, Enum):
    ESEWA = "esewa"
    KHALTI = "khalti"

status = OrderStatus.RECEIVED
print(status.name, status.value)
```

Python `Enum` is a real type; invalid members are harder to construct.

## In other languages

**TypeScript**

```ts
enum OrderStatus { Received, Confirmed, Delivered }
type PayMethod = "esewa" | "khalti" | "cash";  // union — very common
```

**Java** — `enum OrderStatus { RECEIVED, CONFIRMED }` with methods.

**C** — `enum { RECEIVED, CONFIRMED };` (untyped ints unless C++11 `enum class`).

**Rust** — `enum OrderStatus { Received, Confirmed }` — compiler forces you to handle every variant.
