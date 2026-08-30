# 10 — Maps

**Code:** `10_map/maps.go`

## Definition

A map is a hash table: **key → value**. Keys must be comparable (`string`, `int`, structs of comparable fields — not slices). A nil map can be read but **cannot be written** until you `make` it or use a literal.

## Why use

Lookups by name or id: user ages, config flags, counts, caches, grouping. Average get/set is O(1).

## Advantages

- Fast lookup, add, delete.
- Missing key returns the **zero value** (safe, but easy to miss).
- `value, ok := m[key]` tells you if the key exists.
- `range` walks keys/values; `delete`, `clear` (Go 1.21+), `len`.

## Disadvantages

- Nil map panics on write (`m["x"] = 1`).
- Missing key looks like “found 0” unless you use `ok`.
- **Not safe for concurrent write** (use a mutex).
- Range order is randomized — do not rely on it.
- Cannot compare two maps with `==` (only to `nil`).

## How to do it in Go

```go
var nilMap map[string]int     // nil — read ok, write panics

m := make(map[string]string)
m["name"] = "golang"

ages := map[string]int{"alice": 24, "bob": 30}

v, ok := ages["alice"]        // ok idiom
ages["carol"] = 22            // add
ages["alice"] = 25            // update
delete(ages, "bob")
clear(ages)                   // Go 1.21+

for name, age := range ages {
	fmt.Println(name, age)
}
```

## In Python

```python
ages = {"alice": 24, "bob": 30}
ages["carol"] = 22
print(ages.get("carol"))      # None if missing
print(ages["carol"])          # KeyError if missing
del ages["bob"]
ages.clear()

for name, age in ages.items():
    print(name, age)
```

Python `dict` is ordered (3.7+). Missing key with `[]` raises; `.get` is like zero/default.

## In other languages

**JavaScript**

```js
const ages = { alice: 24, bob: 30 };
ages.carol = 22;
delete ages.bob;
Object.keys(ages);
// Map is better for non-string keys:
const m = new Map([["alice", 24]]);
```

**Java** — `HashMap<String, Integer>`, `getOrDefault`, `containsKey`.

**TypeScript** — `Record<string, number>` or `Map<string, number>`.
