# 16 — Structs

**Code:** `16_structs/structs.go`

## Definition

A struct is a custom type that groups related **fields**. Go has no `class`. Methods are functions with a **receiver**: `(o *order)` is like `self` and can change the struct; `(o order)` works on a copy.

There is no constructor keyword. You write `newOrder(...)` yourself.

## Why use

Model real things: orders, users, payments. Fields stay together; methods keep behavior next to the data. Zero values fill missing fields (`""`, `0`, `false`).

## Advantages

- Clear data layout, no hidden class magic.
- Value or pointer receiver — you choose copy vs mutate.
- Literals by field name; anonymous structs for one-off data.
- Comparable if all fields are comparable (`==`).
- Embedding (lesson 17) for composition.

## Disadvantages

- No inheritance (`extends`).
- No real constructor / private `this` setup.
- Forgetting a field silently uses the zero value.
- Large structs copied on value pass — use pointers when mutating or when big.

## How to do it in Go

```go
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

func newOrder(id string, amount float32, status string) *order {
	return &order{id: id, amount: amount, status: status, createdAt: time.Now()}
}

func (o *order) changeStatus(status string) { o.status = status }
func (o order) getAmount() float32          { return o.amount }

o := newOrder("3", 99.5, "pending")
o.changeStatus("shipped")
```

## In Python

```python
from dataclasses import dataclass
from datetime import datetime

@dataclass
class Order:
    id: str
    amount: float
    status: str
    created_at: datetime = None

    def change_status(self, status: str):
        self.status = status

    def get_amount(self) -> float:
        return self.amount
```

Classic `class` + `__init__` is the same idea. Everything is a reference (`self`).

## In other languages

**TypeScript**

```ts
class Order {
  constructor(
    public id: string,
    public amount: number,
    public status: string,
  ) {}
  changeStatus(status: string) { this.status = status; }
}
// or type Order = { id: string; amount: number }
```

**Java** — `class Order { private String id; }` + constructor + getters.

**C** — `struct` with no methods (functions take a pointer).

**Rust** — `struct` + `impl Order { fn change_status(&mut self) }`.
