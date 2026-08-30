package main

import (
	"fmt"
	"sync"
	"time"
)

// ========== WHAT IS A CHANNEL? ==========
// A channel is a typed pipe between goroutines.
// One goroutine SENDS a value in. Another RECEIVES it out.
//
//   ch := make(chan int)   // unbuffered — send waits for a receive
//   ch := make(chan int, 3) // buffered  — send waits only when full
//
//   ch <- 10               // send  (blocks if nobody is ready / buffer full)
//   v := <-ch              // receive (blocks if nothing is there)
//   v, ok := <-ch          // ok is false after the channel is closed and empty
//   close(ch)              // no more sends; receivers still drain leftover values
//
// Go proverb: "Do not communicate by sharing memory;
//              share memory by communicating."
// Channels replace many mutexes: the value is handed off, not shared.
//
// Python: queue.Queue / asyncio.Queue
// JS:     no built-in; use callbacks, Promises, or a tiny queue
// Java:   BlockingQueue
//
// Channel vs WaitGroup:
//   WaitGroup → only wait until N goroutines finish (no data).
//   Channel   → send DATA and/or a "I am done" signal.
//   Many goroutines can share one channel. WaitGroup is not "for multiple only".

// ========== RULES (memorize these) ==========
//  1. Zero value of a channel is nil. Send/receive on nil blocks FOREVER.
//  2. Send on a CLOSED channel PANICS.
//  3. Receive on a closed channel returns the zero value (and ok=false when empty).
//  4. Only the SENDER should close. Never close if other goroutines still send.
//  5. close is not required if the receiver does not range / check closed.
//  6. range ch stops when the channel is closed AND empty.
//  7. Unbuffered send/receive must happen in DIFFERENT goroutines or you deadlock.
//  8. Directional types: chan<- T (send only), <-chan T (receive only).

// ========== UNBUFFERED vs BUFFERED ==========
// Unbuffered (capacity 0): handshake. Sender and receiver meet.
//   send  ────────────►  receive     both wait until the other is ready
//
// Buffered (capacity N): mailbox of size N.
//   send does not block until N values are sitting unread.
//   receive does not block while the mailbox still has values.
//
// Use unbuffered when you want back-pressure / "I know they got it".
// Use buffered for a small queue (emails, jobs) so producers do not stall
// on every item.

// ========== MULTIPLE CHANNELS ==========
// select waits on SEVERAL channels at once (like switch, but for chans).
//   whichever case is ready first, runs.
//   if several are ready, Go picks one at random.
//   default: do this if NONE are ready (non-blocking).
//
// Patterns:
//   select + two result chans     → first API that answers
//   select + time.After           → timeout
//   select + default              → try send/receive without blocking
//   fan-in                        → merge many chans into one
//   fan-out                       → one chan, many worker goroutines
//   pipeline                      → stage1 chan → stage2 chan → stage3

func sum(result chan int, num1, num2 int) {
	result <- num1 + num2 // blocks until main receives
}

func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("processing task")
}

// email worker: range until the job channel is closed, then signal done
func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("sending email to", email)
		time.Sleep(80 * time.Millisecond)
	}
	fmt.Println("email sender: channel closed, all sent")
}

func fakeAPI(name string, ms time.Duration, out chan<- string) {
	time.Sleep(ms)
	out <- name + " ok"
}

// fan-in: copy from many input channels into one output
func merge(chans ...<-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	for _, ch := range chans {
		wg.Add(1)
		go func(c <-chan string) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out) // only closer, after every input is drained
	}()
	return out
}

func produce(name string, n int) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 1; i <= n; i++ {
			ch <- fmt.Sprintf("%s-%d", name, i)
		}
	}()
	return ch
}

// pipeline stage: read in, write squares to out, then close out
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func main() {
	// ----- 1. unbuffered: send and receive must meet -----
	fmt.Println("----- 1. unbuffered result channel -----")
	result := make(chan int)
	go sum(result, 4, 5)
	res := <-result // blocks until sum sends
	fmt.Println("4 + 5 =", res)

	// ----- 2. done channel (instead of Sleep) -----
	fmt.Println("\n----- 2. done signal -----")
	done := make(chan bool)
	go task(done)
	<-done // wait until task finishes
	fmt.Println("task finished")

	// ----- 3. buffered: send N values without a receiver yet -----
	fmt.Println("\n----- 3. buffered channel -----")
	buf := make(chan string, 3)
	buf <- "a" // does not block — room in the mailbox
	buf <- "b"
	buf <- "c"
	// buf <- "d"  // would block forever here (buffer full, no receiver)
	fmt.Println(<-buf, <-buf, <-buf)

	// ----- 4. close + range + comma-ok -----
	fmt.Println("\n----- 4. close, range, comma-ok -----")
	nums := make(chan int, 4)
	nums <- 10
	nums <- 20
	nums <- 30
	close(nums) // no more sends

	v, ok := <-nums
	fmt.Println("first after close still works:", v, "ok?", ok)
	for n := range nums { // stops when closed and empty
		fmt.Println("range leftover:", n)
	}
	v, ok = <-nums
	fmt.Println("empty + closed:", v, "ok?", ok) // 0, false

	// ----- 5. email queue (your example) — producer + worker -----
	fmt.Println("\n----- 5. email sender (buffered jobs + done) -----")
	emailChan := make(chan string, 100)
	emailDone := make(chan bool)
	go emailSender(emailChan, emailDone)
	for i := 0; i < 5; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}
	close(emailChan) // tells range to stop after the last email
	<-emailDone

	// ----- 6. MULTIPLE CHANNELS: select (first ready wins) -----
	fmt.Println("\n----- 6. select — two APIs, take whoever finishes -----")
	users := make(chan string)
	orders := make(chan string)
	go fakeAPI("users", 120*time.Millisecond, users)
	go fakeAPI("orders", 40*time.Millisecond, orders)

	select {
	case u := <-users:
		fmt.Println("select got users first:", u)
	case o := <-orders:
		fmt.Println("select got orders first:", o)
	}

	// drain the slower one so it does not leak
	select {
	case u := <-users:
		fmt.Println("later:", u)
	case o := <-orders:
		fmt.Println("later:", o)
	}

	// ----- 7. select + timeout -----
	fmt.Println("\n----- 7. select + timeout -----")
	slow := make(chan string, 1) // buffer 1 so the late send does not leak a goroutine
	go fakeAPI("reports", 200*time.Millisecond, slow)
	select {
	case r := <-slow:
		fmt.Println("got", r)
	case <-time.After(80 * time.Millisecond):
		fmt.Println("timeout: reports took too long")
	}

	// ----- 8. select + default (non-blocking) -----
	fmt.Println("\n----- 8. non-blocking default -----")
	idle := make(chan int)
	select {
	case v := <-idle:
		fmt.Println("got", v)
	default:
		fmt.Println("nothing ready — do other work")
	}

	// try send without blocking
	full := make(chan int) // unbuffered, no receiver
	select {
	case full <- 1:
		fmt.Println("sent")
	default:
		fmt.Println("cannot send yet — skip")
	}

	// ----- 9. wait on THREE channels -----
	fmt.Println("\n----- 9. select among three channels -----")
	chA := make(chan string)
	chB := make(chan string)
	chC := make(chan string)
	go fakeAPI("sms", 30*time.Millisecond, chA)
	go fakeAPI("push", 50*time.Millisecond, chB)
	go fakeAPI("email", 70*time.Millisecond, chC)
	for i := 0; i < 3; i++ {
		select {
		case v := <-chA:
			fmt.Println(" ", v)
		case v := <-chB:
			fmt.Println(" ", v)
		case v := <-chC:
			fmt.Println(" ", v)
		}
	}

	// ----- 10. fan-in: merge several producers into one channel -----
	fmt.Println("\n----- 10. fan-in (merge multiple channels) -----")
	merged := merge(produce("esewa", 2), produce("khalti", 2), produce("bank", 2))
	for msg := range merged {
		fmt.Println(" ", msg)
	}

	// ----- 11. fan-out: one job channel, many workers -----
	fmt.Println("\n----- 11. fan-out (one channel, 3 workers) -----")
	jobs := make(chan int, 6)
	results := make(chan int, 6)
	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := range jobs {
				fmt.Println("  worker", id, "job", j)
				results <- j * 2
			}
		}(w)
	}
	for j := 1; j <= 6; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
	close(results)
	fmt.Print("  results:")
	for r := range results {
		fmt.Print(" ", r)
	}
	fmt.Println()

	// ----- 12. pipeline: gen → square → print -----
	fmt.Println("\n----- 12. pipeline (multiple channels in a line) -----")
	for n := range square(gen(1, 2, 3, 4)) {
		fmt.Println(" ", n)
	}

	// ----- 13. directional types in the signature -----
	// emailSender(emailChan <-chan string, done chan<- bool)
	// caller cannot send on emailChan inside the worker, or receive on done there.
	fmt.Println("\n----- 13. directional channels -----")
	fmt.Println("  <-chan T  receive-only   chan<- T  send-only")

	// ----- 14. what NOT to do (commented — they panic or deadlock) -----
	// close twice            → panic
	// ch <- x after close    → panic
	// <-make(chan int) in main with no sender → deadlock
	// send on nil channel    → block forever
	fmt.Println("\n----- 14. safety -----")
	fmt.Println("  send on closed → panic | receive on closed → zero, ok=false")
	fmt.Println("  nil channel send/receive → block forever")

	fmt.Println("\nmain finished")
}
