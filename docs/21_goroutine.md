# 21 — Goroutines

**Code:** `21_goroutine/goroutine.go`

## Definition

A goroutine is a function that runs **concurrently** with others. Start it with `go`:

```go
go task(1)
go func() { fmt.Println("hi") }()
```

`main` is itself a goroutine. When `main` returns, the process exits and **all other goroutines are killed**. Prefer `sync.WaitGroup` (or channels) over `time.Sleep`.

A goroutine starts with a ~2 KB stack (grows as needed). An OS thread is typically ~1 MB.

## Concurrency vs parallelism

These are not the same.

| | Concurrency | Parallelism |
|---|-------------|-------------|
| Meaning | Many tasks **in progress** (taking turns) | Many tasks **at the same instant** |
| Needs | Works on **1 CPU** | Needs **2+ cores** and `GOMAXPROCS > 1` |
| Best for | I/O: HTTP, DB, `Sleep` | CPU: image, crypto, tight loops |
| Picture | One chef, two pans | Two chefs, two pans |

Rob Pike: *Concurrency is dealing with lots of things at once. Parallelism is doing lots of things at once.*

**How it performs (G, M, P)**

- **G** — goroutine (`go fn()`)
- **M** — OS thread
- **P** — logical processor; there are `GOMAXPROCS` of them (usually `NumCPU`)

The scheduler is **M:N**: many Gs on fewer Ms. If a G blocks (sleep, channel, syscall), the M can run another G. Idle Ps **steal** work so cores stay busy.

Typical numbers from this lesson on a multi-core machine:

- 5 × 200ms I/O: sequential ~1s, concurrent ~200ms
- CPU work: 1-core goroutines ≈ sequential (or slower); all cores much faster

## Why use

- Call several APIs at once (fan-out), then join.
- `net/http` already runs each request in a goroutine.
- Background jobs: email, thumbnails.
- Worker pools to cap parallelism.
- Pipelines and timeouts.

Not a good fit: tiny sequential math, or work that **must** stay strictly ordered and never waits.

## Advantages

- Cheap to create; simple `go` syntax.
- Scales to huge I/O (servers, crawlers).
- Extra cores used without writing pthreads.
- Blocking one G does not freeze the others.
- Pairs with channels: share memory by communicating.

## Disadvantages

- `main` exit kills them — always wait.
- No guaranteed order (prints interleave).
- Shared variables **race** — use mutex, atomic, or channels. `go run -race`.
- Panic in a goroutine can crash the process.
- Too many CPU goroutines thrash the scheduler.
- Leaked (blocked forever) goroutines waste memory.
- Loop variable capture on Go &lt; 1.22 — pass `i` as an argument.

## How to do it in Go

```go
var wg sync.WaitGroup
for i := 1; i <= 5; i++ {
	wg.Add(1)
	go func(id int) {
		defer wg.Done()
		task(id)
	}(i)
}
wg.Wait()
```

The lesson also shows I/O vs CPU timing, a racy counter vs mutex/atomic, fan-out fake APIs, and a worker pool.

```bash
go run 21_goroutine/goroutine.go
go run -race 21_goroutine/goroutine.go
```

## In Python

**Threads** (I/O; GIL limits CPU parallelism):

```python
import threading

def task(id):
    print("doing task", id)

threads = []
for i in range(6):
    t = threading.Thread(target=task, args=(i,))
    t.start()
    threads.append(t)
for t in threads:
    t.join()
```

**asyncio** (concurrency on one thread — like many waiting tasks):

```python
import asyncio

async def slow(id):
    await asyncio.sleep(0.2)
    print("finished", id)

async def main():
    await asyncio.gather(*(slow(i) for i in range(5)))

asyncio.run(main())
```

**multiprocessing** — real parallelism for CPU (separate processes).

## In other languages

**JavaScript** — one thread + event loop. `Promise.all` is concurrent I/O, not CPU parallelism. Web Workers / Node `worker_threads` for extra cores.

```js
await Promise.all([fetch(a), fetch(b), fetch(c)]);
```

**Java** — `new Thread(runnable).start()`, `ExecutorService`, virtual threads (Loom) — closer to goroutines.

**Rust** — `async` + Tokio (concurrency); `std::thread` or rayon (parallelism).

**C#** — `Task.Run`, `async/await`.
