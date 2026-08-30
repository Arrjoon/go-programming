# 06 — Switch

**Code:** `06_switch/main.go`

## Definition

`switch` picks one branch from many options. In Go, cases **do not fall through** by default (unlike C/Java). You can list several values in one `case`, or write a `switch` with no value (same as if/else if).

## Why use

Cleaner than a long `if / else if` chain when you match one value (day number, HTTP method, enum status). Also used later for **type switches** on interfaces.

## Advantages

- No accidental fall-through (safer than C).
- Several values in one case: `case "a", "e", "i"`.
- Tagless `switch { }` replaces stacked ifs.
- `default` handles the rest.

## Disadvantages

- Fall-through is opt-in (`fallthrough`) and easy to misuse.
- Cases must be constants or comparable values (not arbitrary ranges unless you use a tagless switch).
- Order of cases matters for tagless switches.

## How to do it in Go

```go
switch day {
case 1:
	fmt.Println("Monday")
case 3:
	fmt.Println("Wednesday")
default:
	fmt.Println("Weekend")
}

switch letter {
case "a", "e", "i", "o", "u":
	fmt.Println("vowel")
default:
	fmt.Println("consonant")
}

switch {
case hour < 12:
	fmt.Println("Good morning")
case hour < 18:
	fmt.Println("Good afternoon")
default:
	fmt.Println("Good evening")
}
```

## In Python

Python 3.10+ `match`:

```python
match day:
    case 1:
        print("Monday")
    case 3:
        print("Wednesday")
    case _:
        print("Weekend")

match letter:
    case "a" | "e" | "i" | "o" | "u":
        print("vowel")
    case _:
        print("consonant")
```

Older Python: use `if / elif`. There is no classic C-style `switch`.

## In other languages

**JavaScript**

```js
switch (day) {
  case 1:
    console.log("Monday");
    break; // needed or it falls through
  default:
    console.log("Weekend");
}
```

**Java** — same fall-through rules as JS/C unless you use `->` switch expressions (Java 14+).

**TypeScript** — same as JS, plus enums work well with `switch`.
