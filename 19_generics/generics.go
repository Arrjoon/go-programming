package main

import "fmt"

// ========== WHAT ARE GENERICS? ==========
// Generics let you write ONE function or type that works for MANY types,
// without copying the function and without losing type safety (unlike any).
//
// Without generics you write the same loop twice:
//   printIntSlice([]int)
//   printStringSlice([]string)
//
// With generics you write it once:
//   printSlice[T any](items []T)
//
//   T     = type parameter (a placeholder for a real type)
//   any   = constraint (which types T is allowed to be)
//   [T any] sits between the name and the parameters
//
// TypeScript:
//   function printSlice<T>(items: T[]): void
//   class Stack<T> { push(v: T) {} }
//
// Python:
//   def print_slice[T](items: list[T]) -> None: ...
//
// Go (1.18+):
//   func printSlice[T any](items []T)
//   type stack[T any] struct { items []T }
//
// Why not just use any?
//   func first(items []any) any     ← caller must type-assert the result
//   func first[T any](items []T) T  ← result IS T, compiler checks it
//
// Rules:
//  1. Type params go in [ ] after the name: func name[T constraint](...)
//  2. Constraint = which operations T may use. any means "any type".
//  3. The compiler often INFERS T from the arguments — you can omit [int].
//  4. Use generics for containers, helpers, and repeated algorithms.
//  5. Do NOT use generics when a plain interface or a concrete type is clearer.

// ========== 1. GENERIC FUNCTION (replaces one-type copies) ==========
// Old way: one function per type.
func printStringSlice(items []string) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// Generic way: T is filled in at the call site (int, string, bool, ...).
func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// first element — return type is T, not any
func first[T any](items []T) (T, bool) {
	var zero T
	if len(items) == 0 {
		return zero, false
	}
	return items[0], true
}

func last[T any](items []T) (T, bool) {
	var zero T
	if len(items) == 0 {
		return zero, false
	}
	return items[len(items)-1], true
}

// ========== 2. comparable CONSTRAINT ==========
// comparable = types that support == and !=
//   ints, strings, bools, structs of comparable fields
//   NOT slices, maps, or functions
// Use when: search, unique, map keys, "does this exist?"

func contains[T comparable](items []T, target T) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}
	return false
}

func indexOf[T comparable](items []T, target T) int {
	for i, item := range items {
		if item == target {
			return i
		}
	}
	return -1
}

// ========== 3. UNION / NUMERIC CONSTRAINT ==========
// T must be one of these types. ~int also accepts type Age int
// (the ~ means "this type or any type whose underlying type is int").
//
// Future case: money, scores, stats, conversions.

type Number interface {
	~int | ~int64 | ~float32 | ~float64
}

func sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

func min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type Age int // ~int lets Age work with Number

// ========== 4. MULTIPLE TYPE PARAMETERS ==========
// Future case: maps, pair/tuple, convert A → B, cache key/value.

func keys[K comparable, V any](m map[K]V) []K {
	out := make([]K, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}

func values[K comparable, V any](m map[K]V) []V {
	out := make([]V, 0, len(m))
	for _, v := range m {
		out = append(out, v)
	}
	return out
}

type pair[A any, B any] struct {
	first  A
	second B
}

func makePair[A any, B any](a A, b B) pair[A, B] {
	return pair[A, B]{first: a, second: b}
}

// ========== 5. SLICE HELPERS (map / filter / reduce) ==========
// Future case: any list transform — orders, users, prices.

func mapSlice[T any, U any](items []T, fn func(T) U) []U {
	out := make([]U, len(items))
	for i, item := range items {
		out[i] = fn(item)
	}
	return out
}

func filter[T any](items []T, keep func(T) bool) []T {
	var out []T
	for _, item := range items {
		if keep(item) {
			out = append(out, item)
		}
	}
	return out
}

func reduce[T any, U any](items []T, start U, fn func(U, T) U) U {
	acc := start
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

// ========== 6. GENERIC STRUCT — STACK ==========
// Future case: undo history, parser, navigation, nested menus.

type stack[T any] struct {
	items []T
}

func newStack[T any]() *stack[T] {
	return &stack[T]{}
}

func (s *stack[T]) push(v T) {
	s.items = append(s.items, v)
}

func (s *stack[T]) pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

func (s *stack[T]) peek() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

func (s *stack[T]) len() int {
	return len(s.items)
}

// ========== 7. GENERIC STRUCT — QUEUE ==========
// Future case: job workers, message buffers, print queues.

type queue[T any] struct {
	items []T
}

func newQueue[T any]() *queue[T] {
	return &queue[T]{}
}

func (q *queue[T]) enqueue(v T) {
	q.items = append(q.items, v)
}

func (q *queue[T]) dequeue() (T, bool) {
	var zero T
	if len(q.items) == 0 {
		return zero, false
	}
	v := q.items[0]
	q.items = q.items[1:]
	return v, true
}

// ========== 8. GENERIC SET ==========
// Future case: unique tags, user IDs, visited pages.
// K must be comparable because it is a map key.

type set[T comparable] struct {
	m map[T]struct{}
}

func newSet[T comparable]() *set[T] {
	return &set[T]{m: make(map[T]struct{})}
}

func (s *set[T]) add(v T) {
	s.m[v] = struct{}{}
}

func (s *set[T]) has(v T) bool {
	_, ok := s.m[v]
	return ok
}

func (s *set[T]) remove(v T) {
	delete(s.m, v)
}

func (s *set[T]) size() int {
	return len(s.m)
}

// ========== 9. RESULT[T] — value or error without repeating types ==========
// Future case: APIs, parsers, anything that can fail.

type result[T any] struct {
	value T
	err   error
}

func okResult[T any](v T) result[T] {
	return result[T]{value: v}
}

func errResult[T any](err error) result[T] {
	return result[T]{err: err}
}

func (r result[T]) unwrap() (T, error) {
	return r.value, r.err
}

// ========== 10. GENERIC INTERFACE ==========
// Future case: store/repo that works for User, Order, Product, ...

type store[T any] interface {
	save(item T)
	getAll() []T
}

type memoryStore[T any] struct {
	items []T
}

func (s *memoryStore[T]) save(item T) {
	s.items = append(s.items, item)
}

func (s *memoryStore[T]) getAll() []T {
	return s.items
}

func dumpStore[T any](s store[T]) {
	fmt.Println("store items:", s.getAll())
}

type user struct {
	name string
}

type order struct {
	id     string
	amount float32
}

// ========== 11. POINTER CONSTRAINT (rarely needed) ==========
// *T — useful when you must call methods on a pointer to a new T.
// Most day-to-day code does not need this.

func zeroOf[T any]() T {
	var v T
	return v
}

func main() {
	// ----- 1. one function, many slice types -----
	nums := []int{1, 2, 3, 4, 5}
	names := []string{"golang", "typescript"}
	flags := []bool{true, false, true}

	fmt.Println("ints:")
	printSlice(nums) // compiler infers T = int
	fmt.Println("strings:")
	printSlice[string](names) // you MAY write T yourself
	fmt.Println("bools:")
	printSlice(flags)

	// old one-type function still works, but you do not need it
	printStringSlice([]string{"only", "strings"})

	// ----- 2. typed return (no assertion) -----
	if n, ok := first(nums); ok {
		fmt.Println("first int:", n)
	}
	if s, ok := first(names); ok {
		fmt.Println("first string:", s)
	}
	if _, ok := first([]int{}); !ok {
		fmt.Println("first of empty: none")
	}
	if n, ok := last(nums); ok {
		fmt.Println("last int:", n)
	}

	// ----- 3. comparable: contains / index -----
	fmt.Println("has 3?", contains(nums, 3))
	fmt.Println("has 9?", contains(nums, 9))
	fmt.Println("has golang?", contains(names, "golang"))
	fmt.Println("index of typescript:", indexOf(names, "typescript"))
	fmt.Println("index of rust:", indexOf(names, "rust"))

	// ----- 4. numbers + ~int -----
	fmt.Println("sum ints:", sum(nums))
	fmt.Println("sum floats:", sum([]float64{1.5, 2.5, 3}))
	fmt.Println("min:", min(10, 3), "max:", max(10, 3))
	fmt.Println("sum ages:", sum([]Age{20, 25, 30}))

	// ----- 5. two type params -----
	ages := map[string]int{"arjun": 25, "maya": 22}
	fmt.Println("keys:", keys(ages))
	fmt.Println("values:", values(ages))

	p := makePair("order", 100)
	fmt.Println("pair:", p.first, p.second)

	coord := makePair(27.7, 85.3)
	fmt.Println("coord:", coord.first, coord.second)

	// ----- 6. map / filter / reduce -----
	doubled := mapSlice(nums, func(n int) int { return n * 2 })
	fmt.Println("doubled:", doubled)

	labels := mapSlice(nums, func(n int) string {
		return fmt.Sprintf("#%d", n)
	})
	fmt.Println("labels:", labels)

	evens := filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println("evens:", evens)

	total := reduce(nums, 0, func(acc, n int) int { return acc + n })
	fmt.Println("reduced sum:", total)

	joined := reduce(names, "", func(acc, s string) string {
		if acc == "" {
			return s
		}
		return acc + "," + s
	})
	fmt.Println("joined:", joined)

	// ----- 7. generic stack -----
	intStack := newStack[int]()
	intStack.push(10)
	intStack.push(20)
	intStack.push(30)
	fmt.Println("stack len:", intStack.len())
	if v, ok := intStack.peek(); ok {
		fmt.Println("peek:", v)
	}
	if v, ok := intStack.pop(); ok {
		fmt.Println("pop:", v)
	}
	fmt.Println("stack after pop len:", intStack.len())

	undo := newStack[string]()
	undo.push("type hello")
	undo.push("delete word")
	if action, ok := undo.pop(); ok {
		fmt.Println("undo:", action)
	}

	// ----- 8. generic queue -----
	jobs := newQueue[string]()
	jobs.enqueue("resize image")
	jobs.enqueue("send email")
	if job, ok := jobs.dequeue(); ok {
		fmt.Println("next job:", job)
	}

	// ----- 9. generic set -----
	tags := newSet[string]()
	tags.add("go")
	tags.add("go") // duplicate ignored
	tags.add("web")
	fmt.Println("has go?", tags.has("go"))
	fmt.Println("has rust?", tags.has("rust"))
	fmt.Println("tag count:", tags.size())
	tags.remove("web")
	fmt.Println("after remove:", tags.size())

	ids := newSet[int]()
	ids.add(1)
	ids.add(2)
	fmt.Println("has user 2?", ids.has(2))

	// ----- 10. result[T] -----
	good := okResult(42)
	if v, err := good.unwrap(); err == nil {
		fmt.Println("result ok:", v)
	}
	bad := errResult[int](fmt.Errorf("not found"))
	if _, err := bad.unwrap(); err != nil {
		fmt.Println("result err:", err)
	}

	// ----- 11. generic store / interface -----
	users := &memoryStore[user]{}
	users.save(user{name: "arjun"})
	users.save(user{name: "maya"})
	dumpStore[user](users)

	orders := &memoryStore[order]{}
	orders.save(order{id: "1", amount: 100})
	dumpStore[order](orders)

	// ----- 12. zero value of T -----
	fmt.Println("zero int:", zeroOf[int]())
	fmt.Println("zero string:", fmt.Sprintf("%q", zeroOf[string]()))
	fmt.Println("zero bool:", zeroOf[bool]())
}
