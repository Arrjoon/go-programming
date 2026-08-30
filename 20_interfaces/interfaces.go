package main

import "fmt"

// ========== WHAT IS AN INTERFACE? ==========
// An interface is a CONTRACT: a list of method names + signatures.
// Any type that has those methods AUTOMATICALLY implements the interface.
// There is no "implements" keyword (unlike Java / Python ABC).
//
//   type speaker interface {
//       speak() string
//   }
//
// If dog has speak() string, then dog IS a speaker.
// If cat has speak() string, then cat IS a speaker.
//
// Why use them?
//  1. Write code against the contract, not one concrete type.
//  2. Swap implementations without changing the caller (open/closed).
//  3. Easier unit tests — pass a fake that still has the methods.
//
// Python (ABC / duck typing):
//   class Paymenter(ABC):
//       @abstractmethod
//       def pay(self, amount): ...
//
// Go (implicit):
//   type paymenter interface { pay(amount float32) }
//   // any type with pay(float32) implements paymenter — no extra line
//
// Rules:
//  1. Interface = method set only. No fields.
//  2. Satisfaction is implicit. The type never names the interface.
//  3. Interface values are (type, value). Both can be nil.
//  4. The empty interface (any / interface{}) matches every type.
//  5. Naming: one method often ends in -er (Reader, Writer, paymenter).

// ========== 1. CREATE AN INTERFACE ==========
// Syntax:
//
//	type Name interface {
//	    MethodName(params) returnType
//	}
//
// mainly used in open closed princeple in solid principle
type speaker interface {
	speak() string
}

type dog struct {
	name string
}

func (d dog) speak() string {
	return d.name + " says woof"
}

type cat struct {
	name string
}

func (c cat) speak() string {
	return c.name + " says meow"
}

// announce accepts ANY type that has speak() string
func announce(s speaker) {
	fmt.Println(s.speak())
}

// ========== 2. PAYMENT GATEWAY (open/closed principle) ==========
// Open for addition, closed for modification (SOLID).
// Without an interface you hard-code esewa/razorpay inside makePayment.
// Adding khalti then means editing that function — bad for tests and OCP.
//
// With paymenter, makePayment never changes. You add a new type + inject it.

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay", amount)
}

type esewa struct{}

func (e esewa) pay(amount float32) {
	fmt.Println("payment from esewa", amount)
}

type khalti struct{}

func (k khalti) pay(amount float32) {
	fmt.Println("payment from khalti", amount)
}

// fake for tests — same contract, no real API
type fakeGateway struct {
	lastAmount float32
}

func (f *fakeGateway) pay(amount float32) {
	f.lastAmount = amount
	fmt.Println("fake gateway recorded", amount)
}

// ========== 3. INTERFACE WITH SEVERAL METHODS ==========
type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	w, h float64
}

func (r rectangle) area() float64 {
	return r.w * r.h
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.w + r.h)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func printShape(s shape) {
	fmt.Println("area:", s.area(), "perimeter:", s.perimeter())
}

// ========== 4. EMBED INTERFACES (compose contracts) ==========
// An interface can include other interfaces. The combined type
// must have ALL methods from every embedded interface.
type reader interface {
	read() string
}

type writer interface {
	write(data string)
}

type readWriter interface {
	reader
	writer
}

type notebook struct {
	text string
}

func (n *notebook) read() string {
	return n.text
}

func (n *notebook) write(data string) {
	n.text = data
}

func dump(rw readWriter) {
	rw.write("hello from readWriter")
	fmt.Println("notebook:", rw.read())
}

// ========== 5. EMPTY INTERFACE / any ==========
// interface{} (alias: any) has zero methods, so EVERY type implements it.
// fmt.Println uses this. You lose compile-time type info until you assert.
func printAny(v any) {
	fmt.Println("any value:", v)
}

// ========== 6. TYPE ASSERTION ==========
// i.(T) asks: is the concrete type inside i equal to T?
//
//	val, ok := i.(string)
//
// ok is false if it is not that type (safe).
// i.(string) without ok panics if the type is wrong.
func describe(v any) {
	if s, ok := v.(string); ok {
		fmt.Println("string of length", len(s))
		return
	}
	if n, ok := v.(int); ok {
		fmt.Println("int plus one =", n+1)
		return
	}
	fmt.Println("some other type:", v)
}

// ========== 7. TYPE SWITCH ==========
// switch v := i.(type) checks several concrete types at once.
func typeSwitch(v any) {
	switch val := v.(type) {
	case string:
		fmt.Println("type switch: string", val)
	case int:
		fmt.Println("type switch: int", val)
	case dog:
		fmt.Println("type switch: dog", val.speak())
	case nil:
		fmt.Println("type switch: nil")
	default:
		fmt.Println("type switch: unknown", val)
	}
}

// ========== 8. POINTER vs VALUE RECEIVER ==========
// The method set of T includes methods with value receivers.
// The method set of *T includes both value and pointer receivers.
//
//   func (w walker) walk()   → T and *T both implement walker
//   func (w *runner) run()   → only *T implements runner

type walker interface {
	walk()
}

type person struct {
	name string
}

func (p person) walk() {
	fmt.Println(p.name, "is walking")
}

type runner interface {
	run()
}

type athlete struct {
	name string
}

func (a *athlete) run() {
	fmt.Println(a.name, "is running")
}

// ========== 9. NIL INTERFACE vs NIL INSIDE ==========
// var i speaker          → i is nil (no type, no value)
// var d *dog; i = d      → i is NOT nil (type is *dog, value is nil)
func checkSpeaker(s speaker) {
	if s == nil {
		fmt.Println("interface is nil (no type, no value)")
		return
	}
	fmt.Printf("interface is not nil — dynamic type %T\n", s)
	// do not call s.speak() if the inner pointer is nil — that panics
}

func main() {
	// ----- 1. implicit implementation -----
	announce(dog{name: "bruno"})
	announce(cat{name: "milo"})

	// store different types in one interface slice
	pets := []speaker{
		dog{name: "rocky"},
		cat{name: "luna"},
	}
	fmt.Println("all pets:")
	for _, p := range pets {
		fmt.Println(" ", p.speak())
	}

	// ----- 2. inject the gateway — caller chooses the type -----
	p1 := payment{gateway: razorpay{}}
	p1.makePayment(100)

	p2 := payment{gateway: esewa{}}
	p2.makePayment(250)

	p3 := payment{gateway: khalti{}}
	p3.makePayment(75)

	fake := &fakeGateway{}
	p4 := payment{gateway: fake}
	p4.makePayment(10)
	fmt.Println("fake last amount:", fake.lastAmount)

	// ----- 3. one function, two shapes -----
	printShape(rectangle{w: 4, h: 5})
	printShape(circle{radius: 3})

	shapes := []shape{
		rectangle{w: 2, h: 3},
		circle{radius: 1},
	}
	fmt.Println("shapes:")
	for _, s := range shapes {
		printShape(s)
	}

	// ----- 4. embedded interfaces -----
	nb := &notebook{}
	dump(nb)

	// ----- 5. empty interface / any -----
	printAny("arjun")
	printAny(25)
	printAny(true)
	printAny(dog{name: "max"})

	values := []any{"hi", 42, 3.14}
	fmt.Println("slice of any:", values)

	// ----- 6. type assertion -----
	describe("golang")
	describe(7)
	describe(true)

	var i any = "hello"
	s := i.(string)
	fmt.Println("assert without ok:", s)
	// n := i.(int)  // would panic — i is a string, not int

	// ----- 7. type switch -----
	typeSwitch("text")
	typeSwitch(99)
	typeSwitch(dog{name: "zeus"})
	typeSwitch(nil)
	typeSwitch(3.14)

	// ----- 8. value vs pointer receiver -----
	p := person{name: "maya"}
	var w walker = p // value receiver — value is fine
	w.walk()
	w = &p // *person also works for value-receiver methods
	w.walk()

	a := athlete{name: "hari"}
	var r runner = &a // pointer receiver — must pass *athlete
	r.run()
	// var r2 runner = a  // would not compile — athlete has no run(), *athlete does

	// ----- 9. nil checks -----
	checkSpeaker(nil)

	var d *dog
	checkSpeaker(d) // type is *dog, value is nil — interface itself is not nil
}
