# 11 — Range

**Code:** `11_range/11_range.go`

## Definition

`range` is the iterator form of `for`. It walks a slice, array, map, string, channel, or (Go 1.22+) an integer. You get an index/key and a value; use `_` to ignore either.

On a **string**, the value is a **rune** (Unicode code point), and the index is a **byte** offset — important for Nepali and other non-ASCII text.

## Why use

The default way to loop collections. Safer than `for i := 0; i < len; i++` for maps (no index) and strings (multi-byte characters).

## Advantages

- Works on slice, array, map, string, channel, int.
- `_` skips what you do not need.
- Unicode-aware on strings (full characters, not broken bytes).
- Empty/nil slice: body never runs (no crash).

## Disadvantages

- Map order is random.
- String index is in **bytes**, so it jumps (नेपाल is not 0,1,2,3,4).
- `range` on a slice copies each element (large structs: range over index or pointers).
- Old Go loop-variable capture with goroutines (fixed in 1.22).

## How to do it in Go

```go
for i, v := range nums { }     // index + value
for i := range nums { }        // index only
for _, v := range nums { }     // value only
for range nums { count++ }     // just count

for name, age := range ages { }
for name := range ages { }

for i, r := range "नेपाल" {
	fmt.Println(i, string(r))
}

for i := range 5 { }           // 0..4  (Go 1.22+)
```

## In Python

```python
for i, v in enumerate(nums):
    print(i, v)
for v in nums:
    print(v)
for name, age in ages.items():
    print(name, age)
for ch in "नेपाल":             # each character
    print(ch)
for i in range(5):
    print(i)
```

Python `for x in collection` is the closest idea. `enumerate` = index + value.

## In other languages

**JavaScript**

```js
for (const [i, v] of nums.entries()) {}
for (const v of nums) {}
for (const [k, v] of Object.entries(ages)) {}
for (const ch of "नेपाल") {}
```

**Java** — enhanced for-each; `IntStream.range(0, 5)`; strings need `codePoints()` for runes.

**C#** — `foreach (var x in list)`.
