# 24 — Mutex

**Code:** `24_mutex/mutex.go`

## Definition

A **mutex** (MUTual EXclusion) is a lock. Only **one** goroutine at a time may run the code between `Lock` and `Unlock`. That block is the **critical section** — the lines that read or write **shared** data.

```go
var mu sync.Mutex
mu.Lock()
views++          // only one goroutine here at a time
mu.Unlock()
```

Always `Unlock` in a `defer` so a panic still releases the lock.

`count++` is **not atomic**: it is read → add → write. Two goroutines can both read `10` and both write `11`. You expected `12`. That is a **data race**.

## Why use

Several goroutines are changing the **same** variable, map, slice, or struct field.

| Shared thing | What goes wrong without a lock |
|---|---|
| `views++` / `balance += n` | Lost updates |
| `map[string]T` write | Race, often a **panic** |
| `append` on one slice | Corrupt slice header |
| if-not-exists-then-set | Two goroutines both insert |

Go proverb still applies: *share memory by communicating* (channels). Use a mutex when the design **is** shared memory — one post, one wallet, one cache.

## Where to use

- **Post / video view counters** (the `post` example in this lesson)
- Wallet, cart, inventory totals
- In-memory **cache** (`map[string]User`)
- Shared config: many reads, occasional write → `RWMutex`
- Rate-limit counters, in-process connection tables

## Where not to use

- Each goroutine owns its own data — no lock needed
- You only wait until work finishes → `WaitGroup`
- You pass results on a **channel** — no shared write
- A single `int64` increment → `sync/atomic` is enough and faster

## Advantages

- Correct shared updates
- Works for any type (not only ints)
- `RWMutex` allows many readers together
- Simple: lock, change, unlock

## Disadvantages

- Forgetting `Lock` means the mutex does nothing (the old `inc` only called `Unlock`)
- `Unlock` without `Lock` **panics**
- Lock twice on the same `Mutex` **deadlocks** (it is not reentrant)
- Copying a struct that contains a mutex = two locks, one piece of data
- Heavy contention makes everyone wait
- Pipelines are often clearer with channels

## How to do it in Go

**Post views (this repo)**

```go
type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer wg.Done()
	p.mu.Lock()
	defer p.mu.Unlock()
	p.views++
}
```

Receiver must be `*post`. If you pass `post` by value you **copy the mutex** and the lock no longer protects the original `views`.

**Map cache — `RWMutex`**

```go
c.mu.Lock()      // Set — exclusive
c.data[k] = v
c.mu.Unlock()

c.mu.RLock()     // Get — many readers
v, ok := c.data[k]
c.mu.RUnlock()
```

**Mutex vs atomic vs channel**

| Tool | Use when |
|---|---|
| `sync.Mutex` | Any shared struct / map / slice |
| `sync.RWMutex` | Lots of reads, few writes |
| `atomic.AddInt64` | One counter, nothing else in the section |
| channel | Hand off a value; avoid shared writes |

```bash
go run 24_mutex/mutex.go
go run -race 24_mutex/mutex.go
```

## Rules

1. Pair `Lock` / `Unlock`. Prefer `defer Unlock()`.
2. Keep the locked region **small** (do I/O **outside** the lock).
3. Do not copy a mutex.
4. One lock order if you ever take two mutexes (prevents deadlock).
5. `TryLock` (Go 1.18+): return false if the lock is busy — skip or retry.

## In Python

```python
import threading

lock = threading.Lock()
views = 0

def inc():
    global views
    with lock:          # Lock / Unlock
        views += 1

threads = [threading.Thread(target=inc) for _ in range(100)]
for t in threads:
    t.start()
for t in threads:
    t.join()
print(views)            # 100
```

`threading.RLock` is reentrant (Go’s `Mutex` is not). `asyncio` uses a single thread — use `asyncio.Lock` only among coroutines.

## In other languages

**Java**

```java
synchronized (this) { views++; }
// or ReentrantLock lock = new ReentrantLock();
```

**JavaScript** — one thread for normal JS. Shared `Atomics` / `Worker` only when you use workers.

**Rust** — `Mutex<T>`; the compiler forces you to lock to touch `T`.

**C#** — `lock (obj) { views++; }` or `Mutex` / `ReaderWriterLockSlim`.
