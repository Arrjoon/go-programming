# 12 — Functions

**Code:** `12_functions/functions.go`

## Definition

A function is a named block you can call. In Go you declare it at package level (or as a literal). Functions can take parameters, return one or more values, be stored in variables, passed as arguments, or returned from other functions.

## Why use

Reuse logic (greet, add, divide), split `main` into pieces, and return `(value, error)` — the standard Go error style. First-class functions enable callbacks, map/filter, and closures (lesson 14).

## Advantages

- Multiple return values (no wrapper object required).
- Named returns and bare `return`.
- Functions are values (`type mathFunc func(int, int) int`).
- `defer` runs cleanup when the function returns.
- Unused results must be assigned or `_` — you notice ignored errors.

## Disadvantages

- No default arguments and no keyword arguments.
- No function overloading (same name, different types).
- Named returns can hide bugs if you forget to set them.
- Variadic `...` only on the **last** parameter (see lesson 13).

## How to do it in Go

```go
func greet() { fmt.Println("hello") }
func add(a, b int) int { return a + b }
func intro(name string, age int) string { return name }

func getLanguages() (string, string, string) {
	return "golang", "javascript", "c"
}

func divide(a, b float64) (quot float64, ok bool) {
	if b == 0 {
		return 0, false
	}
	quot = a / b
	ok = true
	return
}

lang1, lang2, _ := getLanguages()
```

## In Python

```python
def greet():
    print("hello")

def add(a: int, b: int) -> int:
    return a + b

def get_languages():
    return "golang", "javascript", "c"   # tuple unpack

lang1, lang2, _ = get_languages()

def divide(a, b):
    if b == 0:
        return 0, False
    return a / b, True
```

Python allows defaults (`b=0`) and `*args` / `**kwargs`.

## In other languages

**JavaScript**

```js
function add(a, b) { return a + b; }
const add = (a, b) => a + b;
```

Returns one value (or an array/object). Errors via `throw` or `Result` libraries.

**Java** — `int add(int a, int b)`; overloading allowed; no multiple returns (use a class or record).

**TypeScript** — typed params and returns: `function add(a: number, b: number): number`.
