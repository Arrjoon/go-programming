# 09 — Slices

**Code:** `09_slices/slices.go`

## Definition

A slice is a **dynamic view** over an array: pointer + length + capacity. It can grow with `append`. A nil slice has `len` 0 and equals `nil`. This is Go’s everyday list type (like Python `list`).

- **len** — how many items you have now
- **cap** — how many items the backing array can hold before `append` allocates a bigger one

## Why use

Almost always prefer slices over arrays: shopping lists, query results, tokens, bytes of a file. APIs in the standard library take `[]T`, not `[N]T`.

## Advantages

- Grows with `append` — no manual realloc.
- Cheap to pass (the header is 3 words; data is shared).
- Slicing `s[1:3]` is a view, not a full copy.
- `make([]int, len, cap)` lets you pre-size for speed.

## Disadvantages

- Sharing a backing array: changing one slice can change another.
- `append` may allocate a new array — old slices still see the old data.
- Nil vs empty (`[]int{}`) both have len 0 but only nil equals `nil`.
- Accidental `append` on a shared slice can overwrite later items.

## How to do it in Go

```go
var abc []int                 // nil slice
fmt.Println(abc == nil)       // true
fmt.Println(len(abc))         // 0

nepal := make([]int, 2)       // len 2, cap 2, values 0, 0
nums := make([]int, 2, 5)     // len 2, cap 5
nums = append(nums, 2, 3, 4)

fmt.Println(len(nums), cap(nums))

// literal
names := []string{"go", "ts"}
```

When `append` exceeds `cap`, Go allocates a larger array (often ~2×) and copies.

## In Python

```python
abc = []                      # empty list
nepal = [0, 0]
nums = [0, 0]
nums.append(2)
nums.extend([3, 4])
print(len(nums))              # no separate "capacity"
```

Python lists resize automatically. There is no `cap` you manage.

## In other languages

**JavaScript**

```js
const nums = [];
nums.push(2, 3, 4);
nums.length;
```

**Java** — `ArrayList<Integer> nums = new ArrayList<>();` `add`, `size`, optional `ensureCapacity`.

**C++** — `std::vector<int>` (`size` / `capacity`) is the closest match to a Go slice.
