# 04 — Constants

**Code:** `04_constants/main.go`

## Definition

A constant is a name for a value that is fixed at **compile time**. You cannot assign to it later. Go also has `iota` — an auto-incrementing counter inside a `const` block (0, 1, 2, …).

## Why use

Use constants for values that must not change: π, HTTP status codes, days in a week, config flags. `iota` is the usual way to start **enums** (see lesson 18).

## Advantages

- Compiler blocks accidental changes (`pi = 3` will not compile).
- Untyped constants are more flexible in expressions.
- `iota` avoids writing 0, 1, 2 by hand.
- Grouped `const (` blocks stay readable.

## Disadvantages

- Only compile-time values (no `const now = time.Now()`).
- Slices, maps, and structs from runtime data cannot be `const`.
- `iota` resets in every new `const` block — easy to forget.

## How to do it in Go

```go
const pi = 3.14159
const greeting = "Hello"
const daysInWeek int = 7

const (
	statusOK    = 200
	statusError = 500
)

const (
	Sunday = iota // 0
	Monday        // 1
	Tuesday       // 2
)
```

## In Python

```python
PI = 3.14159          # convention only — still assignable
GREETING = "Hello"

from enum import IntEnum
class Weekday(IntEnum):
    SUNDAY = 0
    MONDAY = 1
    TUESDAY = 2
```

Python has no real `const`. `Final` in `typing` is a hint, not enforced at runtime (unless you use a linter).

## In other languages

**JavaScript**

```js
const pi = 3.14159;   // cannot reassign the binding
```

Objects inside `const` can still change (`obj.x = 1`).

**Java**

```java
final double PI = 3.14159;
static final int STATUS_OK = 200;
```

**TypeScript**

```ts
const daysInWeek = 7;
enum Weekday { Sunday, Monday, Tuesday }
```
