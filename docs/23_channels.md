# 23 — Channels

**Code:** `23_channels/channels.go`

## Definition

A **channel** is a typed pipe between goroutines. One side **sends** a value in; another **receives** it out. The channel carries both **data** and **synchronization** (the send waits until a receive is ready, unless there is buffer space).

```go
ch := make(chan int)      // unbuffered
ch := make(chan int, 3)   // buffered, capacity 3

ch <- 10                  // send
v := <-ch                 // receive
v, ok := <-ch             // ok == false when closed and empty
close(ch)
```

Go proverb: *Do not communicate by sharing memory; share memory by communicating.*

## Why use

- Hand a result from a worker back to `main` (`sum` → `result`).
- Signal “I am done” (`done <- true`) instead of `time.Sleep`.
- Queue jobs (emails, thumbnails) with a buffered channel + `range`.
- Wait on **several** sources at once with `select` (APIs, timeout, quit).
- Fan-in / fan-out / pipelines — the usual concurrent designs.

**Channel vs WaitGroup**

| | Channel | WaitGroup |
|---|---|---|
| Sends data? | Yes | No — only a count |
| “Finished”? | Receive on `done`, or close + range | `Wait()` |
| Many goroutines? | Yes — they can all send on one chan | Yes — that is its job |

Use **WaitGroup** when you only need to wait. Use a **channel** when you need values or a quit/done signal. They often appear together (workers + `jobs` chan + `wg.Wait()`).

## Advantages

- Type-safe handoff (`chan string` cannot send an `int`).
- Back-pressure: unbuffered send waits until someone is listening.
- `select` waits on many channels without busy-looping.
- `close` + `range` is a clean “no more work” signal.
- Directional types (`<-chan`, `chan<-`) document who may send.
- Avoids many shared-memory races (the value is copied through the pipe).

## Disadvantages

- Send on a **closed** channel **panics**.
- Receive on **nil** (or send on nil) **blocks forever**.
- Unbuffered send/receive in the **same** goroutine **deadlocks**.
- Forgetting `close` makes `range` hang forever.
- Closing when another goroutine still sends → panic.
- Extra goroutines + channels are harder to debug than straight-line code.
- Buffered size is a guess: too small stalls producers; too large hides slowness.

## Rules

1. Zero value is `nil` — send/receive on it never completes.
2. Only the **sender** closes. Do not close if others still send.
3. After `close`, leftover values can still be received; then `ok` is `false`.
4. `range ch` ends when the channel is closed **and** empty.
5. Unbuffered = handshake (both sides ready). Buffered = mailbox of size N.
6. `select` picks a ready case; if several are ready, one is chosen **at random**.
7. `default` runs when **no** case is ready (non-blocking try).

## How channels work (unbuffered vs buffered)

**Unbuffered** (`make(chan int)`): send and receive happen together. The sender blocks until a receiver is waiting (and the other way around). Good when you need to know the value was taken.

**Buffered** (`make(chan int, 100)`): the sender blocks only when the buffer is **full**. The receiver blocks only when the buffer is **empty**. Good for a small job queue (the email example).

```
unbuffered:   sender ──── handshake ──── receiver
buffered:     sender ──► [ | | | ] ──► receiver
```

## Multiple channels

`select` is how you work with **more than one** channel at the same time.

```go
select {
case u := <-users:
    fmt.Println("users", u)
case o := <-orders:
    fmt.Println("orders", o)
case <-time.After(80 * time.Millisecond):
    fmt.Println("timeout")
default:
    fmt.Println("nothing ready")
}
```

| Pattern | What it does | In the lesson |
|---|---|---|
| `select` two results | First API that answers | example 6 |
| `select` + `time.After` | Give up if too slow | example 7 |
| `select` + `default` | Try send/receive without blocking | example 8 |
| Three channels in a loop | Collect sms, push, email as they finish | example 9 |
| **Fan-in** | Merge many producers → one `out` | `merge(...)` |
| **Fan-out** | One `jobs` chan, many workers | 3 workers, 6 jobs |
| **Pipeline** | `gen → square → print` | chained channels |

Fan-in: each input is ranged in its own goroutine; a WaitGroup closes `out` when all inputs are done.

Fan-out: workers all `range jobs`. Close `jobs` after enqueueing so they exit.

## How to do it in Go

**Result + done**

```go
result := make(chan int)
go func() { result <- 4 + 5 }()
fmt.Println(<-result)

done := make(chan bool)
go func() {
    defer func() { done <- true }()
    fmt.Println("work")
}()
<-done
```

**Email queue (this repo)**

```go
emailChan := make(chan string, 100)
done := make(chan bool)
go emailSender(emailChan, done)
for i := 0; i < 5; i++ {
    emailChan <- fmt.Sprintf("%d@gmail.com", i)
}
close(emailChan)
<-done
```

`emailSender` uses `for email := range emailChan` so it stops only after `close`.

**Directional**

```go
func emailSender(emailChan <-chan string, done chan<- bool)
```

The worker may only receive emails and only send on `done`.

```bash
go run 23_channels/channels.go
```

## What not to do

```go
// ch <- 1 after close(ch)     → panic
// close(ch); close(ch)        → panic
// <-make(chan int) in main    → deadlock (no sender)
// var ch chan int; ch <- 1    → block forever (nil)
```

## In Python

```python
from queue import Queue
from threading import Thread

q = Queue(maxsize=100)

def email_sender():
    while True:
        email = q.get()
        if email is None:      # like close
            q.task_done()
            break
        print("sending", email)
        q.task_done()

t = Thread(target=email_sender)
t.start()
for i in range(5):
    q.put(f"{i}@gmail.com")
q.put(None)
q.join()
t.join()
```

`asyncio.Queue` is the async version. Python has no `select` on queues; `asyncio.wait` / `select` on sockets is the cousin.

## In other languages

**Java** — `BlockingQueue`, `LinkedBlockingQueue`. `take()` / `put()` block like unbuffered/buffered channels.

**JavaScript** — no built-in channel. `Promise.race([fetchA, fetchB, timeout])` is like `select` + timeout.

```js
const winner = await Promise.race([
  fetch("/users").then((r) => r.json()),
  new Promise((_, reject) => setTimeout(() => reject("timeout"), 80)),
]);
```

**Rust** — `std::sync::mpsc` (`tx.send`, `rx.recv`); `tokio::select!`.

**C#** — `Channel<T>` in `System.Threading.Channels`; `WaitAny` on several waits.
