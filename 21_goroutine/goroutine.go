package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// ========== WHAT IS A GOROUTINE? ==========
// A goroutine is a function that runs CONCURRENTLY with other functions.
// Start one with the go keyword:
//
//   go task(1)          ← returns immediately; task runs in the background
//   go func() { ... }() ← anonymous goroutine
//
// main itself is a goroutine. When main returns, the PROCESS exits and
// every other goroutine is killed — even if they are still working.
// That is why beginners use time.Sleep. Prefer WaitGroup or channels.
//
// Compared to an OS thread:
//   OS thread  → ~1 MB stack, created by the kernel, expensive
//   goroutine  → ~2 KB stack (grows as needed), created by the Go runtime
//
// You can start hundreds of thousands of goroutines. You cannot do that
// with OS threads.
//
// Python: threading.Thread / asyncio.create_task
// JS:     Promise / setTimeout (one thread + event loop)
// Go:     go fn()  — runtime multiplexes goroutines onto OS threads

// ========== CONCURRENCY vs PARALLELISM ==========
// These words are not the same.
//
// CONCURRENCY = structure. Many tasks IN PROGRESS. You deal with them
// at once by switching. One chef, two pans: stir A, then B, then A.
// Works even on ONE CPU. Goroutines take turns (the scheduler preempts).
//
// PARALLELISM = execution. Many tasks AT THE SAME INSTANT. Two chefs,
// two pans, same moment. Needs TWO OR MORE CPU cores and GOMAXPROCS > 1.
//
//   1 core:   G1 ---- G2 ---- G1 ---- G2     concurrent, not parallel
//   4 cores:  G1 ====================         concurrent AND parallel
//             G2 ====================
//             G3 ====================
//
// Rob Pike: "Concurrency is about dealing with lots of things at once.
//            Parallelism is about doing lots of things at once."
//
// Go gives you concurrency by default (go keyword).
// Parallelism happens when the scheduler puts goroutines on different
// OS threads that the OS runs on different cores.

// ========== HOW THE RUNTIME PERFORMS (G, M, P) ==========
// Go uses an M:N scheduler — many goroutines (G) on fewer OS threads (M).
//
//   G  goroutine     your go fn()
//   M  machine       a real OS thread
//   P  processor     a logical CPU slot (there are GOMAXPROCS of these)
//
// A P holds a run queue of Gs. An M must hold a P to execute a G.
//
//   GOMAXPROCS = how many Gs can run in PARALLEL (usually = NumCPU)
//   runtime.GOMAXPROCS(1) → only concurrency (one core of Go work)
//   runtime.GOMAXPROCS(8) → up to 8 goroutines truly in parallel
//
// What happens when a goroutine blocks?
//   time.Sleep / channel wait / mutex  → scheduler parks that G,
//                                        M picks another G from the queue
//   syscall (network, disk)            → M may detach; another M + P
//                                        keeps running other Gs
//
// Work stealing: an idle P steals Gs from a busy P's queue so cores
// stay busy. That is why fan-out jobs finish faster on multi-core.
//
// I/O-bound work (HTTP, DB, sleep):
//   Thousands of goroutines sit waiting. Almost no CPU. Concurrency wins.
//   100 API calls of 200ms → ~200ms total, not 20 seconds.
//
// CPU-bound work (image, crypto, tight loops):
//   Extra goroutines only help if they run on extra cores (parallelism).
//   On 1 core, 4 CPU goroutines ≈ same time as 1 (plus a little overhead).
//   On 4 cores, 4 CPU goroutines ≈ 4x faster (ideally).
//
// Tiny tasks: starting a goroutine is cheap but not free. Do not spawn
// one goroutine per 1+1 addition. Batch or use a worker pool.

// ========== ADVANTAGES ==========
//  1. Cheap to create — ~2 KB, not 1 MB like a thread.
//  2. Simple syntax — one word: go.
//  3. Scales to huge I/O (HTTP servers, crawlers, chat).
//  4. Scheduler handles blocking so other work continues.
//  5. Same code can use extra cores without you writing pthreads.
//  6. Pairs with channels for "share memory by communicating".

// ========== DISADVANTAGES / GOTCHAS ==========
//  1. main exit kills them. Always wait (WaitGroup / channel / context).
//  2. No guaranteed order. Prints will interleave.
//  3. Shared variables race. Use mutex, atomic, or channels.
//     go run -race file.go detects this.
//  4. Loop variable capture (Go < 1.22): all goroutines saw the last i.
//     Go 1.22+ gives each iteration its own i. Still safer to pass i as arg.
//  5. Panic in a goroutine can crash the whole program (use recover if needed).
//  6. Too many CPU goroutines thrash the scheduler (use NumCPU workers).
//  7. Harder to debug than straight-line code (stack traces per G).
//  8. Leaked goroutines (blocked forever) waste memory — always have an exit.

// ========== USE CASES ==========
//  1. Independent I/O: call several APIs / DBs at once, then join results.
//  2. HTTP server: each request already runs in its own goroutine (net/http).
//  3. Background work: emails, thumbnails, cleanup, after a request returns.
//  4. Worker pool: N workers pull jobs from a channel (rate-limit CPU/I/O).
//  5. Timeouts / heartbeats: a goroutine sleeps, then cancels or pings.
//  6. Pipeline: stage1 → stage2 → stage3, each stage a set of goroutines.
//  7. Fan-out / fan-in: split work, then merge results.
//  8. Not a good fit: tiny sequential math, or work that MUST be ordered
//     and has no waiting.

func task(id int) {
	fmt.Println("doing task", id)
}

func slowTask(id int, ms time.Duration) {
	time.Sleep(ms) // stand-in for HTTP / DB / disk
	fmt.Println("finished slow task", id)
}

// ========== sequential vs concurrent I/O ==========
func runSequentialIO(n int, each time.Duration) time.Duration {
	start := time.Now()
	for i := 1; i <= n; i++ {
		time.Sleep(each)
	}
	return time.Since(start)
}

func runConcurrentIO(n int, each time.Duration) time.Duration {
	start := time.Now()
	var wg sync.WaitGroup
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(each)
		}()
	}
	wg.Wait()
	return time.Since(start)
}

// ========== CPU-bound work (parallelism demo) ==========
func cpuWork(iters int) int {
	x := 0
	for i := 0; i < iters; i++ {
		x += i % 7
	}
	return x
}

func runSequentialCPU(workers, iters int) time.Duration {
	start := time.Now()
	for i := 0; i < workers; i++ {
		_ = cpuWork(iters)
	}
	return time.Since(start)
}

func runParallelCPU(workers, iters int) time.Duration {
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = cpuWork(iters)
		}()
	}
	wg.Wait()
	return time.Since(start)
}

// ========== shared counter: race vs safe ==========
func raceCounter(n int) int {
	count := 0
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++ // DATA RACE — many goroutines write the same int
		}()
	}
	wg.Wait()
	return count
}

func mutexCounter(n int) int {
	count := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}
	wg.Wait()
	return count
}

func atomicCounter(n int) int64 {
	var count int64
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&count, 1)
		}()
	}
	wg.Wait()
	return count
}

// ========== fan-out: fake API calls, collect results ==========
func fakeAPI(name string, ms time.Duration) string {
	time.Sleep(ms)
	return name + " ok"
}

func fetchAll(names []string, delay time.Duration) []string {
	results := make([]string, len(names))
	var wg sync.WaitGroup
	for i, name := range names {
		wg.Add(1)
		go func(i int, name string) {
			defer wg.Done()
			results[i] = fakeAPI(name, delay) // each index is unique — no race
		}(i, name)
	}
	wg.Wait()
	return results
}

// ========== worker pool ==========
func workerPool(jobs []int, workers int) []int {
	jobCh := make(chan int)
	out := make([]int, len(jobs))
	var wg sync.WaitGroup

	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range jobCh {
				out[i] = jobs[i] * jobs[i] // "process" job i
			}
		}()
	}

	for i := range jobs {
		jobCh <- i
	}
	close(jobCh)
	wg.Wait()
	return out
}

func main() {
	fmt.Println("CPUs:", runtime.NumCPU(), "GOMAXPROCS:", runtime.GOMAXPROCS(0))

	// ----- 1. fire and forget — BAD wait: Sleep -----
	// tasks may finish in ANY order. Sleep is a guess, not a guarantee.
	fmt.Println("\n----- 1. go task (Sleep wait) -----")
	for i := 0; i <= 5; i++ {
		go task(i)
	}
	time.Sleep(100 * time.Millisecond)

	// ----- 2. correct wait: WaitGroup -----
	// Add(1) before go, Done() when finished, Wait() in main.
	fmt.Println("\n----- 2. WaitGroup -----")
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			task(id)
		}(i) // pass i in — safe on every Go version
	}
	wg.Wait()
	fmt.Println("all WaitGroup tasks done")

	// ----- 3. anonymous goroutine -----
	fmt.Println("\n----- 3. anonymous go func -----")
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("hello from anonymous goroutine")
	}()
	wg.Wait()

	// ----- 4. CONCURRENCY (I/O): 5 x 200ms -----
	// sequential ≈ 1000ms   concurrent ≈ 200ms
	// One CPU is enough — they mostly SLEEP, not compute.
	fmt.Println("\n----- 4. concurrency: I/O-bound -----")
	seqIO := runSequentialIO(5, 200*time.Millisecond)
	conIO := runConcurrentIO(5, 200*time.Millisecond)
	fmt.Println("sequential 5x200ms:", seqIO)
	fmt.Println("concurrent 5x200ms:", conIO)
	fmt.Println("concurrency saved about", seqIO-conIO)

	// ----- 5. PARALLELISM (CPU): same work, 1 core vs all cores -----
	workers := runtime.NumCPU()
	if workers < 2 {
		workers = 2
	}
	iters := 25_000_000

	fmt.Println("\n----- 5. parallelism: CPU-bound -----")
	old := runtime.GOMAXPROCS(1)
	oneCore := runParallelCPU(workers, iters)
	runtime.GOMAXPROCS(old) // back to all cores
	manyCore := runParallelCPU(workers, iters)
	seqCPU := runSequentialCPU(workers, iters)

	fmt.Println("workers:", workers, "iters each:", iters)
	fmt.Println("sequential (no goroutines):", seqCPU)
	fmt.Println("goroutines on 1 core  (concurrent, not parallel):", oneCore)
	fmt.Println("goroutines on all cores (concurrent + parallel):", manyCore)
	fmt.Println("1-core ≈ sequential (switching cost). all-cores should be faster.")

	// ----- 6. shared memory without lock — WRONG count -----
	// run: go run -race goroutine.go   to see the race detector
	fmt.Println("\n----- 6. race vs mutex vs atomic -----")
	n := 1000
	fmt.Println("expected:", n)
	fmt.Println("racy count (often < 1000):", raceCounter(n))
	fmt.Println("mutex count:", mutexCounter(n))
	fmt.Println("atomic count:", atomicCounter(n))

	// ----- 7. fan-out I/O (use case: many APIs at once) -----
	fmt.Println("\n----- 7. fan-out fake APIs -----")
	start := time.Now()
	got := fetchAll([]string{"users", "orders", "payments", "stock"}, 150*time.Millisecond)
	fmt.Println("results:", got, "in", time.Since(start), "(~150ms, not 600ms)")

	// ----- 8. worker pool (use case: limit parallelism) -----
	fmt.Println("\n----- 8. worker pool -----")
	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("squares:", workerPool(jobs, 3))

	// ----- 9. background slow tasks with WaitGroup -----
	fmt.Println("\n----- 9. several slow tasks -----")
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			slowTask(id, 80*time.Millisecond)
		}(i)
	}
	wg.Wait()

	fmt.Println("\nmain finished — all waited goroutines completed")
}
