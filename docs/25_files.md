# 25 — Files

**Code:** `25_fiels/files.go`  
(folder name is `25_fiels`; the program is `files.go`)

## Definition

File I/O is reading and writing **bytes on disk**. In Go the handle is `*os.File`. You **open** (or **create**) a file, **read** or **write**, then **close** it.

```go
sourceFile, err := os.Open("example.txt")
if err != nil {
	panic(err)
}
defer sourceFile.Close()
```

`defer Close()` runs when the function returns — even after a panic — so the OS file handle is always released.

## Why use

- Load config, logs, CSV, images
- Save user uploads or generated reports
- Copy / backup / rename / delete
- List a folder (`ReadDir`)
- Stream a **large** file in chunks (buffer) so you do not load it all into RAM

## Advantages

- `defer Close()` is hard to forget once you use it
- `os.ReadFile` / `os.WriteFile` are one-liners for small files
- `io.Copy` copies any reader to any writer (files, network, buffers)
- `bufio` adds fast buffered read/write
- Errors are values — you always check `err`

## Disadvantages

- Forgetting `Close()` leaks handles
- `ReadFile` on a huge file can use a lot of memory
- `os.Create` **truncates** an existing file (wipes it)
- Paths depend on the **current working directory**
- Concurrent writes to one file need a lock or one writer

## How to do it in Go

Run from the lesson folder so `example.txt` is found:

```bash
cd 25_fiels
go run files.go
```

| Task | API |
|---|---|
| Open read-only | `os.Open(name)` |
| Create / overwrite | `os.Create(name)` |
| Append | `os.OpenFile(name, os.O_APPEND\|os.O_WRONLY, 0644)` |
| Info | `f.Stat()` — name, size, time, is-dir |
| Chunk read | `f.Read(buf)` |
| Whole file | `os.ReadFile` / `os.WriteFile` |
| Copy | `io.Copy(dst, src)` or `bufio` byte loop |
| Folder | `os.Open(dir)` then `ReadDir(-1)` |
| Exists? | `os.Stat` + `os.IsNotExist(err)` |
| Rename | `os.Rename` |
| Delete | `os.Remove` |

**Buffer** = a `[]byte` in RAM. `Read` fills it from disk. You do not load the whole file unless you use `ReadFile`.

**EOF** = end of file. It is not a failure. Check with `errors.Is(err, io.EOF)`, not `err.Error() == "EOF"`.

**Flush:** `bufio.Writer` keeps data in memory until `Flush()` (or `Close`). Call `Flush()` after a copy loop.

**Permissions:** `0644` means owner read/write, others read. On Windows the mode is mostly ignored.

## Examples in the lesson

1. Open + `Stat`
2. Read 12 bytes into a buffer
3. `ReadFile` — entire `example.txt`
4. `Create` + `WriteString` → `example2.txt`
5. Append a line
6. Your **bufio** copy (`ReadByte` / `WriteByte`)
7. Same copy with **`io.Copy`** (preferred)
8. `WriteFile`
9. Exists check
10. Rename
11. List current folder
12. Delete the temp files (`example.txt` is kept)

## In Python

```python
with open("example.txt") as f:          # close is automatic
    print(f.read())

with open("example2.txt", "w") as f:
    f.write("hello\n")

with open("example2.txt", "a") as f:
    f.write("appended\n")

import shutil
shutil.copyfile("example.txt", "copy.txt")
os.remove("copy.txt")
os.listdir(".")
```

`with` is Python’s `defer Close()`.

## In other languages

**JavaScript (Node)**

```js
import fs from "fs";
const text = fs.readFileSync("example.txt", "utf8");
fs.writeFileSync("example2.txt", "hello\n");
fs.copyFileSync("example.txt", "copy.txt");
fs.unlinkSync("copy.txt");
```

**Java** — `Files.readString(path)`, `Files.writeString`, `Files.copy`, `Files.delete`.

**C** — `fopen` / `fread` / `fwrite` / `fclose` (you must close yourself, like Go without `defer`).
