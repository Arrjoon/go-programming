# 01 — Hello World

**Code:** `01_hello_world/main.go`

## Definition

A Hello World program is the smallest complete program: it starts, prints a message, and exits. In Go every runnable program needs a `package main` and a `func main()`.

## Why use

- Confirms the Go toolchain works (`go run`).
- Shows the required entry point (`main`).
- Introduces `fmt` — the standard way to print.

## Advantages

- Fast to compile and run (no VM startup like Java).
- One binary later (`go build`) — no interpreter required.
- Simple, fixed program structure.

## Disadvantages

- More ceremony than a Python one-liner (`print("hi")`).
- You must declare `package` and `main` even for a tiny script.

## How to do it in Go

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Welcome to Go programming!")
}
```

```bash
go run 01_hello_world/main.go
```

## In Python

```python
print("Hello, World!")
print("Welcome to Go programming!")
```

No `main` is required. The file runs from top to bottom.

## In other languages

**JavaScript (Node)**

```js
console.log("Hello, World!");
```

**Java** — class + `public static void main`:

```java
public class Main {
    public static void main(String[] args) {
        System.out.println("Hello, World!");
    }
}
```

**C**

```c
#include <stdio.h>
int main() {
    printf("Hello, World!\n");
    return 0;
}
```

Go sits between C and Python: compiled like C, but with garbage collection and a simpler syntax than Java.
