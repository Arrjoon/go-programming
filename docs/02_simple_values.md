# 02 — Simple values

**Code:** `02_simple_values/main.go`

## Definition

Simple values (also called literals or primitive values) are the basic data you work with: strings, integers, floats, and booleans. You can print them and combine them with operators (`+`, `/`, `&&`, `||`, `!`).

## Why use

Every program stores and computes data. These types are the building blocks for variables, structs, and APIs later.

## Advantages

- Types are checked at compile time (unlike Python).
- Operators are familiar (`+` for strings and numbers).
- Booleans are real `true` / `false` (not 1 / 0).

## Disadvantages

- Integer division: `7 / 3` is `2` (int). Use `7.0 / 3.0` for a float.
- You cannot mix types freely (`int + float64` needs a conversion).
- Strings are immutable — concatenation creates a new string.

## How to do it in Go

```go
fmt.Println("go" + "lang")     // string concat
fmt.Println("1 + 1 =", 1+1)    // int
fmt.Println("7.0 / 3.0 =", 7.0/3.0)
fmt.Println(true && false)     // AND
fmt.Println(true || false)     // OR
fmt.Println(!true)             // NOT
```

Common types: `string`, `int`, `int64`, `float32`, `float64`, `bool`, `byte`, `rune`.

## In Python

```python
print("go" + "lang")
print("1 + 1 =", 1 + 1)
print("7.0 / 3.0 =", 7.0 / 3.0)
print(True and False)
print(True or False)
print(not True)
```

Python uses `and` / `or` / `not` instead of `&&` / `||` / `!`. Types are dynamic.

## In other languages

**JavaScript**

```js
console.log("go" + "lang");
console.log(1 + 1);
console.log(7 / 3);          // always float-like number
console.log(true && false);
```

**Java** — must declare types: `String`, `int`, `double`, `boolean`.

**TypeScript** — same operators as JS, plus static types: `string`, `number`, `boolean`.
