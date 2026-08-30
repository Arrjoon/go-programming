package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// ========== WHAT IS A MUTEX? ==========
// Mutex = MUTual EXclusion. A lock so only ONE goroutine at a time
// can run the critical section (the lines that touch shared data).
//
//   var mu sync.Mutex
//   mu.Lock()     // wait here if someone else holds the lock
//   // ... change the shared variable ...
//   mu.Unlock()   // let the next waiter in
//
// Without a lock, two goroutines can read the same value, both add 1,
// both write back — you lose an update. That is a DATA RACE.
//
//   views is 10
//   G1 reads 10          G2 reads 10
//   G1 writes 11         G2 writes 11     ← expected 12, got 11
//
// Go proverb: "Do not communicate by sharing memory;
//              share memory by communicating."
// Prefer a channel when you can pass the data. Use a mutex when many
// goroutines must UPDATE THE SAME variable / map / struct field.
//
// Python: threading.Lock()
// JS:     single-threaded — no mutex for normal code
// Java:   synchronized / ReentrantLock

// ========== WHY USE ==========
//  1. count++ / views++ / balance += n is NOT atomic (read + write).
//  2. Maps are NOT safe for concurrent write (can panic).
//  3. Slices: append from many goroutines corrupts the header.
//  4. "Check then update" (if !exists { m[k] = v }) must be one lock.
//
// WHERE to use:
//   - view / like counters on a post
//   - wallet / cart totals
//   - in-memory cache or map[string]User
//   - shared config that goroutines read and sometimes write
//   - rate-limit counters, connection pools
//
// WHERE NOT to use:
//   - each goroutine has its own data (no sharing)
//   - you only wait for finish → WaitGroup
//   - you hand results through a channel → no shared write
//   - a single int increment → sync/atomic is enough (and faster)

// ========== RULES ==========
//  1. Unlock in a defer so panics still release the lock.
//  2. Lock and Unlock must be paired. Unlock without Lock panics.
//  3. Do NOT copy a mutex (or a struct that contains one) by value.
//     Pass *post, not post. The copy has a different lock.
//  4. Keep the locked region SMALL — only the shared read/write.
//  5. Same lock order if you ever take two mutexes (avoids deadlock).
//  6. RWMutex: many readers OR one writer — good for read-heavy maps.
//  7. Detect races: go run -race mutex.go

// ========== ADVANTAGES ==========
//  - Correct shared updates (no lost increment).
//  - Simple mental model: lock, change, unlock.
//  - Works for any data (int, map, struct, slice).
//  - RWMutex lets many readers proceed together.

// ========== DISADVANTAGES ==========
//  - Easy to forget Lock (your old inc only Unlocked — lock did nothing).
//  - Deadlock if you Lock twice on the same mutex (it is not reentrant).
//  - Forgotten Unlock blocks every other goroutine forever.
//  - Copied mutex = fake safety (two locks, one piece of data).
//  - Contended locks slow the program (everyone queues).
//  - Does not compose as cleanly as channels for pipelines.

type post struct {
	views int
	mu    sync.Mutex // do not copy post by value — mutex must stay unique
}

// inc — YOUR example, fixed: Lock before changing views
func (p *post) inc(wg *sync.WaitGroup) {
	defer wg.Done()
	p.mu.Lock()
	defer p.mu.Unlock() // unlock even if we panic later
	p.views++
}

func (p *post) viewCount() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.views
}

// same idea without a lock — WRONG when many goroutines call it
func (p *post) incUnsafe() {
	p.views++
}

// ========== bank account (another everyday case) ==========
type account struct {
	mu      sync.Mutex
	balance int
}

func (a *account) deposit(n int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += n
}

func (a *account) withdraw(n int) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.balance < n {
		return false
	}
	a.balance -= n
	return true
}

func (a *account) get() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

// ========== map cache — maps panic if written from two goroutines ==========
type cache struct {
	mu   sync.RWMutex // many Get, few Set
	data map[string]string
}

func newCache() *cache {
	return &cache{data: make(map[string]string)}
}

func (c *cache) set(k, v string) {
	c.mu.Lock() // exclusive — writers wait for each other
	defer c.mu.Unlock()
	c.data[k] = v
}

func (c *cache) get(k string) (string, bool) {
	c.mu.RLock() // shared — many readers at once
	defer c.mu.RUnlock()
	v, ok := c.data[k]
	return v, ok
}

func (c *cache) size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.data)
}

func raceCounter(n int) int {
	count := 0
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++ // DATA RACE
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

func atomicCount(n int) int64 {
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

func main() {
	// ----- 1. YOUR case: post views from 100 goroutines -----
	fmt.Println("----- 1. post views (mutex) -----")
	var wg sync.WaitGroup
	myPost := post{views: 0}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}
	wg.Wait()
	fmt.Println("views (want 100):", myPost.viewCount())

	// ----- 2. same increment WITHOUT a lock — often < 100 -----
	fmt.Println("\n----- 2. post views WITHOUT mutex -----")
	unsafe := post{views: 0}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			unsafe.incUnsafe()
		}()
	}
	wg.Wait()
	fmt.Println("views unsafe (often < 100):", unsafe.views)

	// ----- 3. race vs mutex vs atomic -----
	fmt.Println("\n----- 3. counter 1000 times -----")
	n := 1000
	fmt.Println("expected:", n)
	fmt.Println("racy:   ", raceCounter(n))
	fmt.Println("mutex:  ", mutexCounter(n))
	fmt.Println("atomic: ", atomicCount(n))

	// ----- 4. bank — deposit / withdraw share one balance -----
	fmt.Println("\n----- 4. bank account -----")
	acc := &account{balance: 0}
	for i := 0; i < 50; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			acc.deposit(10)
		}()
		go func() {
			defer wg.Done()
			acc.withdraw(3)
		}()
	}
	wg.Wait()
	fmt.Println("balance (50*10 - successful withdraws):", acc.get())

	// ----- 5. cache map with RWMutex -----
	fmt.Println("\n----- 5. cache (RWMutex) -----")
	c := newCache()
	c.set("user:1", "arjun")
	c.set("user:2", "maya")

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if v, ok := c.get("user:1"); ok {
				fmt.Println("  reader", id, "got", v)
			}
		}(i)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.set("user:3", "hari")
	}()
	wg.Wait()
	fmt.Println("cache size:", c.size())

	// ----- 6. TryLock (Go 1.18+) — do other work if busy -----
	fmt.Println("\n----- 6. TryLock -----")
	var mu sync.Mutex
	mu.Lock()
	if mu.TryLock() {
		fmt.Println("got second lock") // will not print
		mu.Unlock()
	} else {
		fmt.Println("lock busy — skip or retry later")
	}
	mu.Unlock()

	// ----- 7. keep the critical section small -----
	fmt.Println("\n----- 7. lock only the shared write -----")
	var hits int
	var hitsMu sync.Mutex
	start := time.Now()
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(20 * time.Millisecond) // I/O — no lock needed
			hitsMu.Lock()
			hits++
			hitsMu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("hits:", hits, "in", time.Since(start), "(~20ms, not 160ms)")

	// ----- 8. what NOT to do (commented — deadlock / panic) -----
	// mu.Lock(); mu.Lock()     // deadlock — Mutex is NOT reentrant
	// var m sync.Mutex; m.Unlock() // panic — unlock of unlocked mutex
	// p2 := myPost; go p2.inc()    // COPIED mutex — does not protect myPost
	fmt.Println("\n----- 8. safety -----")
	fmt.Println("  always Lock before Unlock")
	fmt.Println("  never copy a struct that contains a Mutex")
	fmt.Println("  run: go run -race 24_mutex/mutex.go")

	fmt.Println("\nmain finished")
}
