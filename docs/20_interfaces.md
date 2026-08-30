# 20 — Interfaces

**Code:** `20_interfaces/interfaces.go`

## Definition

An interface is a **contract**: a list of method signatures. Any type that has those methods **implements** the interface automatically. There is no `implements` keyword.

```go
type paymenter interface {
	pay(amount float32)
}
```

`razorpay`, `esewa`, and `khalti` all implement `paymenter` just by having `pay`.

An interface value is `(dynamic type, dynamic value)`. The empty interface is `any` (`interface{}`) — every type implements it.

## Why use

Write `makePayment` against the contract, not one gateway. Add `khalti` without editing the function (**open/closed**). Pass a `fakeGateway` in tests. Same idea for `shape`, `readWriter`, and stores.

## Advantages

- Implicit satisfaction — less coupling.
- Swap implementations (prod vs fake).
- Small interfaces (`io.Reader`) compose well.
- Embedding interfaces (`readWriter` = reader + writer).
- Type assert / type switch when you need the concrete type.

## Disadvantages

- Implicit impl can surprise you (a method rename silently breaks it).
- `any` throws away type safety until you assert.
- `i.(T)` without `ok` panics.
- Nil interface vs interface holding a nil pointer — `s == nil` can be false, then a method call panics.
- Pointer vs value receivers: `func (a *athlete) run()` means only `*athlete` implements `runner`.

## How to do it in Go

```go
type speaker interface{ speak() string }

func announce(s speaker) { fmt.Println(s.speak()) }
announce(dog{name: "bruno"})

p := payment{gateway: esewa{}}
p.makePayment(100)

if s, ok := v.(string); ok { /* ... */ }
switch val := v.(type) {
case int:
	// ...
}
```

## In Python

```python
from abc import ABC, abstractmethod

class Paymenter(ABC):
    @abstractmethod
    def pay(self, amount: float) -> None: ...

class Esewa(Paymenter):
    def pay(self, amount: float) -> None:
        print("esewa", amount)

# duck typing without ABC: any object with .pay() works
def make_payment(gateway, amount):
    gateway.pay(amount)
```

Python is duck-typed at runtime. `Protocol` (typing) is the closest to Go’s implicit interface for checkers.

## In other languages

**Java / C#**

```java
interface Paymenter { void pay(float amount); }
class Esewa implements Paymenter { public void pay(float amount) {} }
```

You must write `implements`.

**TypeScript**

```ts
interface Paymenter { pay(amount: number): void }
// structural — extra `implements` is optional
```

**Rust** — `trait Paymenter { fn pay(&self, amount: f32); }` + `impl Paymenter for Esewa`.
