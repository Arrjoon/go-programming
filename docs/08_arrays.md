# 08 — Arrays

**Code:** `08_arrays/main.go`

## Definition

An array is a **fixed-size** sequence of values of the **same type**. The size is part of the type: `[5]int` is not the same type as `[3]int`. Index starts at 0. Unset slots are zero values.

## Why use

When the length is known and must not change: RGB (`[3]byte`), a 2D matrix, small buffers. Most day-to-day lists in Go should be **slices** (lesson 09), not arrays.

## Advantages

- Size is known at compile time — no surprise growth.
- Values live in one contiguous block (good for tiny, hot data).
- Zero-filled on create — safe defaults.
- `len`, index, and `range` work like other collections.

## Disadvantages

- Cannot grow. `[5]int` stays length 5.
- Passing an array to a function **copies the whole array**.
- `[3]int` and `[5]int` are different types — awkward APIs.
- Rarely what you want for user lists, JSON, or DB rows.

## How to do it in Go

```go
var numbers [5]int          // [0 0 0 0 0]
numbers[0] = 10

fruits := [3]string{"apple", "banana", "mango"}
scores := [...]int{90, 85, 78, 92, 23} // Go counts the size

for i := 0; i < len(fruits); i++ {
	fmt.Println(i, fruits[i])
}
for index, value := range scores {
	fmt.Println(index, "->", value)
}

matrix := [2][3]int{
	{1, 2, 3},
	{4, 5, 6},
}
```

## In Python

```python
# Python has no fixed array in the language. Use a list:
numbers = [0] * 5
numbers[0] = 10
fruits = ["apple", "banana", "mango"]

# array module (same type, still resizable)
import array
scores = array.array("i", [90, 85, 78])

# numpy for real numeric arrays
import numpy as np
matrix = np.array([[1, 2, 3], [4, 5, 6]])
```

Python lists grow; they are closer to **slices**.

## In other languages

**JavaScript** — `Array` is dynamic (like a slice). Typed arrays exist (`Int32Array`).

**Java**

```java
int[] numbers = new int[5];
String[] fruits = {"apple", "banana", "mango"};
int[][] matrix = {{1, 2, 3}, {4, 5, 6}};
```

Length is fixed after `new`. `ArrayList` is the growable version.

**C** — `int numbers[5];` — closest to Go arrays, including copy-on-pass if you wrap in a struct.
