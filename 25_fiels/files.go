package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

// ========== WHAT IS FILE I/O? ==========
// A file on disk is bytes. In Go you talk to it through *os.File:
//
//   f, err := os.Open("example.txt")   // read only
//   if err != nil { panic(err) }
//   defer f.Close()                    // always close, even if we return/panic
//
// Open  = existing file, read only
// Create = new file (or truncate existing), write
// OpenFile = full control (append, read+write, permissions)
// ReadFile / WriteFile = whole file in one call (small files)
//
// Why defer Close()?
//   The OS has a limit on open files. Close releases the handle and
//   flushes buffered writes. defer runs when the function returns.
//
// Python: open("f") as f / f.read() / f.write()
// JS:     fs.readFileSync / fs.promises
// Java:   FileInputStream / Files.readAllBytes

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// make sure the sample file exists so every example below can run
	must(os.WriteFile("example.txt", []byte("hello golang\n"), 0644))

	// ----- 1. open + stat (name, size, time) -----
	fmt.Println("----- 1. open and stat -----")
	f, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println("name:", info.Name())
	fmt.Println("size (bytes):", info.Size())
	fmt.Println("is dir?", info.IsDir())
	fmt.Println("modified:", info.ModTime())
	fmt.Println("mode:", info.Mode())

	// ----- 2. read into a buffer (chunk of memory) -----
	// a buffer is temporary space in RAM. Read fills it from disk.
	fmt.Println("\n----- 2. read into a buffer -----")
	src, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	buf := make([]byte, 12)
	n, err := src.Read(buf)
	if err != nil && !errors.Is(err, io.EOF) {
		panic(err)
	}
	fmt.Println("bytes read:", n)
	fmt.Println("text:", string(buf[:n]))

	// ----- 3. read the WHOLE file at once (small files only) -----
	fmt.Println("\n----- 3. ReadFile (all bytes) -----")
	all, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(all))

	// ----- 4. create + write -----
	fmt.Println("----- 4. create and write -----")
	out, err := os.Create("example2.txt") // creates or empties the file
	if err != nil {
		panic(err)
	}
	// Close this one ourselves before we copy/delete later
	_, err = out.WriteString("hello this is written by myself\n")
	must(err)
	_, err = out.WriteString("this is the next line\n")
	must(err)
	must(out.Close())
	fmt.Println("wrote example2.txt")

	// ----- 5. append (do not wipe the file) -----
	fmt.Println("\n----- 5. append -----")
	app, err := os.OpenFile("example2.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	_, err = app.WriteString("appended line\n")
	must(err)
	must(app.Close())
	data, _ := os.ReadFile("example2.txt")
	fmt.Println(string(data))

	// ----- 6. copy with bufio (your byte-by-byte example) -----
	fmt.Println("----- 6. copy with bufio -----")
	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create("example_copy.txt")
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) { // end of file — not a real error
				break
			}
			panic(err)
		}
		if err := writer.WriteByte(b); err != nil {
			panic(err)
		}
	}
	must(writer.Flush()) // push leftover buffer bytes to disk
	fmt.Println("copied example.txt -> example_copy.txt (bufio)")

	// ----- 7. copy the clean way: io.Copy -----
	fmt.Println("\n----- 7. copy with io.Copy -----")
	from, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer from.Close()

	to, err := os.Create("example_copy2.txt")
	if err != nil {
		panic(err)
	}
	defer to.Close()

	written, err := io.Copy(to, from)
	if err != nil {
		panic(err)
	}
	fmt.Println("io.Copy bytes:", written)

	// ----- 8. WriteFile — create/overwrite in one line -----
	fmt.Println("\n----- 8. WriteFile -----")
	must(os.WriteFile("notes.txt", []byte("line one\nline two\n"), 0644))
	fmt.Println("wrote notes.txt")

	// ----- 9. check if a file exists -----
	fmt.Println("\n----- 9. exists? -----")
	_, err = os.Stat("example.txt")
	if err == nil {
		fmt.Println("example.txt exists")
	} else if os.IsNotExist(err) {
		fmt.Println("example.txt missing")
	} else {
		panic(err)
	}

	// ----- 10. rename / move -----
	fmt.Println("\n----- 10. rename -----")
	must(os.Rename("notes.txt", "notes_renamed.txt"))
	fmt.Println("notes.txt -> notes_renamed.txt")

	// ----- 11. read a folder -----
	fmt.Println("\n----- 11. read directory -----")
	dir, err := os.Open(".")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	entries, err := dir.ReadDir(-1) // -1 = all entries
	if err != nil {
		panic(err)
	}
	for _, e := range entries {
		kind := "file"
		if e.IsDir() {
			kind = "dir"
		}
		fmt.Println(" ", kind, e.Name())
	}

	// ----- 12. delete -----
	fmt.Println("\n----- 12. delete -----")
	must(os.Remove("example2.txt"))
	must(os.Remove("example_copy.txt"))
	must(os.Remove("example_copy2.txt"))
	must(os.Remove("notes_renamed.txt"))
	fmt.Println("deleted temp files (kept example.txt)")

	fmt.Println("\nmain finished")
}
