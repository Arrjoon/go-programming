package main

import "fmt"

// ========== WHAT IS AN ENUM IN GO? ==========
// Go has NO built-in enum keyword (unlike TypeScript / Java / Python Enum).
// We fake enums with:
//   1. a custom type  (so you cannot mix OrderStatus with a plain int by accident)
//   2. a const block  (the named allowed values)
//
// Why not just use strings?
//   status := "delieverd"   ← typo compiles, bug at runtime
//   changeStatus(Delivered) ← compiler only accepts OrderStatus values
//
// Two main styles:
//   type OrderStatus int     + iota     → small, fast, good for switches
//   type OrderStatus string  + "paid"   → readable in logs, JSON, APIs
//
// TypeScript:
//   enum Status { Received, Confirmed, Delivered }
//   type Status = "received" | "confirmed" | "delivered"
//
// Go:
//   type Status int
//   const ( Received Status = iota; Confirmed; Delivered )
//
// Rules:
//  1. iota resets to 0 in every const () block, then +1 per line.
//  2. The zero value is 0 / "". Design so zero means "unset" or the first state.
//  3. Add String(), IsValid(), and ParseX() when the value comes from users/APIs.
//  4. A typed const is still just a number or string at runtime — validate input.

// ========== 1. INTEGER ENUM (iota) ==========
// Use when: internal state machines, switches, you do not need the name in JSON.
// iota = 0, 1, 2, 3, ... automatically.
//
// Future case: any workflow with a fixed set of steps
//   order, ticket, kyc, shipment, interview, game level, ...

type OrderStatus int

const (
	Received  OrderStatus = iota // 0 — new order just placed
	Confirmed                    // 1
	Prepared                     // 2
	Delivered                    // 3
	Cancelled                    // 4
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status, "=", status.String())
}

// String lets fmt.Println print the name, not 0/1/2.
func (s OrderStatus) String() string {
	switch s {
	case Received:
		return "received"
	case Confirmed:
		return "confirmed"
	case Prepared:
		return "prepared"
	case Delivered:
		return "delivered"
	case Cancelled:
		return "cancelled"
	default:
		return fmt.Sprintf("unknown(%d)", s)
	}
}

func (s OrderStatus) IsValid() bool {
	return s >= Received && s <= Cancelled
}

// ParseOrderStatus turns user/API input into the typed enum.
func ParseOrderStatus(s string) (OrderStatus, bool) {
	switch s {
	case "received":
		return Received, true
	case "confirmed":
		return Confirmed, true
	case "prepared":
		return Prepared, true
	case "delivered":
		return Delivered, true
	case "cancelled":
		return Cancelled, true
	default:
		return 0, false
	}
}

// ========== 2. STRING ENUM ==========
// Use when: logs, HTTP, JSON, DB columns should show "esewa" not 2.
// Safer to read; slightly more memory than int.
//
// Future case: API payloads, payment gateways, locales, currency codes.

type OrderStatusString string

const (
	StatusReceived  OrderStatusString = "received"
	StatusConfirmed OrderStatusString = "confirmed"
	StatusPrepared  OrderStatusString = "prepared"
	StatusDelivered OrderStatusString = "delivered"
)

func changeOrderStatusString(status OrderStatusString) {
	fmt.Println("changing order status to", status)
}

// ========== 3. IOTA TRICKS ==========
// blank identifier skips a number
// start from 1 so 0 can mean "unset"
// expressions work: each line reuses the same formula with a new iota

const (
	_ = iota // skip 0
	one
	two
	three
)

type HTTPStatus int

const (
	StatusUnset HTTPStatus = iota // 0 = not set (zero value)
	StatusOK                      // 1
	StatusCreated                 // 2
	StatusNotFound                // 3
	StatusError                   // 4
)

// file sizes: iota in an expression
const (
	_  = iota             // skip 0
	KB = 1 << (10 * iota) // 1 << 10 = 1024
	MB                    // 1 << 20
	GB                    // 1 << 30
)

// ========== 4. BIT FLAGS (combine several options) ==========
// Use when: one value holds MANY on/off switches (permissions, features).
// Check with & , turn on with | , turn off with &^
//
// Future case: file perms, feature flags, notification channels, UI options.

type Permission uint

const (
	PermRead    Permission = 1 << iota // 1  (1 << 0) — iota starts at 0 here
	PermWrite                          // 2  (1 << 1)
	PermExecute                        // 4  (1 << 2)
	PermAdmin                          // 8  (1 << 3)
)

const PermNone Permission = 0 // keep zero out of the iota block so bits stay 1,2,4,8

func (p Permission) Has(flag Permission) bool {
	return p&flag != 0
}

// ========== 5. USER ROLE ==========
// Future case: auth, middleware, admin panels.
// Keep roles as an enum so "adimn" typos cannot sneak in.

type Role int

const (
	RoleGuest Role = iota
	RoleUser
	RoleModerator
	RoleAdmin
)

func (r Role) String() string {
	switch r {
	case RoleGuest:
		return "guest"
	case RoleUser:
		return "user"
	case RoleModerator:
		return "moderator"
	case RoleAdmin:
		return "admin"
	default:
		return "unknown"
	}
}

func (r Role) CanEdit() bool {
	return r == RoleModerator || r == RoleAdmin
}

// ========== 6. PAYMENT METHOD ==========
// Future case: checkout, invoices, refunds — same idea as interfaces.go gateways.

type PayMethod string

const (
	PayCash    PayMethod = "cash"
	PayCard    PayMethod = "card"
	PayEsewa   PayMethod = "esewa"
	PayKhalti  PayMethod = "khalti"
	PayRazor   PayMethod = "razorpay"
	PayBank    PayMethod = "bank"
)

func (m PayMethod) IsOnline() bool {
	switch m {
	case PayEsewa, PayKhalti, PayRazor, PayCard:
		return true
	default:
		return false
	}
}

// ========== 7. PRIORITY / SEVERITY ==========
// Future case: support tickets, logs, job queues, alerts.
// Integer enum is nice here because you can compare: High > Medium.

type Priority int

const (
	PriorityLow Priority = iota + 1 // start at 1; 0 means unset
	PriorityMedium
	PriorityHigh
	PriorityCritical
)

func (p Priority) String() string {
	switch p {
	case PriorityLow:
		return "low"
	case PriorityMedium:
		return "medium"
	case PriorityHigh:
		return "high"
	case PriorityCritical:
		return "critical"
	default:
		return "unset"
	}
}

type LogLevel int

const (
	LogDebug LogLevel = iota
	LogInfo
	LogWarn
	LogError
	LogFatal
)

// ========== 8. ENVIRONMENT ==========
// Future case: config, feature toggles, which DB/API URL to use.

type Env string

const (
	EnvDev     Env = "dev"
	EnvStaging Env = "staging"
	EnvProd    Env = "prod"
)

func (e Env) IsProd() bool {
	return e == EnvProd
}

// ========== 9. WEEKDAY ==========
// Future case: scheduling, opening hours, cron-like jobs.
// iota matches time.Weekday style (Sunday = 0) if you start there.

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (d Weekday) IsWeekend() bool {
	return d == Saturday || d == Sunday
}

// ========== 10. DIRECTION / STATE MACHINE ==========
// Future case: games, UI navigation, traffic, undo/redo, player movement.

type Direction int

const (
	DirNorth Direction = iota
	DirEast
	DirSouth
	DirWest
)

func (d Direction) TurnRight() Direction {
	return (d + 1) % 4
}

// booking / ticket lifecycle — only some transitions are legal
type TicketState int

const (
	TicketOpen TicketState = iota
	TicketInProgress
	TicketResolved
	TicketClosed
)

func (from TicketState) CanGo(to TicketState) bool {
	switch from {
	case TicketOpen:
		return to == TicketInProgress || to == TicketClosed
	case TicketInProgress:
		return to == TicketResolved || to == TicketOpen
	case TicketResolved:
		return to == TicketClosed || to == TicketInProgress
	default:
		return false // closed is final
	}
}

// ========== 11. OPTIONAL / TRI-STATE ==========
// Future case: filters, form fields, "yes / no / not answered".
// Do not use bool when you need a third "unknown" value.

type TriState int

const (
	TriUnknown TriState = iota
	TriYes
	TriNo
)

// ========== 12. SWITCH ON ENUM ==========
// Always have default — an int can still hold 99 even if that is not a const.

func handleOrder(status OrderStatus) {
	switch status {
	case Received:
		fmt.Println("handle: send confirmation email")
	case Confirmed:
		fmt.Println("handle: start packing")
	case Prepared:
		fmt.Println("handle: hand to rider")
	case Delivered:
		fmt.Println("handle: ask for review")
	case Cancelled:
		fmt.Println("handle: refund")
	default:
		fmt.Println("handle: unknown status", status)
	}
}

func main() {
	// ----- 1. int enum + iota -----
	changeOrderStatus(Received)
	changeOrderStatus(Confirmed)
	changeOrderStatus(Delivered)
	fmt.Println("Received is", int(Received), "Delivered is", int(Delivered))

	// ----- 2. string enum -----
	changeOrderStatusString(StatusReceived)
	changeOrderStatusString(StatusConfirmed)

	// ----- 3. String(), IsValid(), Parse -----
	fmt.Println("name of 2:", Prepared.String())
	fmt.Println("Prepared valid?", Prepared.IsValid())
	fmt.Println("99 valid?", OrderStatus(99).IsValid())

	if s, ok := ParseOrderStatus("delivered"); ok {
		fmt.Println("parsed:", s)
	}
	if _, ok := ParseOrderStatus("delieverd"); !ok {
		fmt.Println("typo rejected: delieverd")
	}

	// ----- 4. switch -----
	handleOrder(Prepared)
	handleOrder(Cancelled)
	handleOrder(OrderStatus(99))

	// ----- 5. iota tricks -----
	fmt.Println("skipped 0 → one,two,three:", one, two, three)
	fmt.Println("HTTP zero means unset:", StatusUnset, StatusOK)
	fmt.Println("sizes KB,MB,GB:", KB, MB, GB)

	// ----- 6. bit flags -----
	userPerms := PermRead | PermWrite
	fmt.Println("user perms:", userPerms)
	fmt.Println("can read?", userPerms.Has(PermRead))
	fmt.Println("can exec?", userPerms.Has(PermExecute))
	adminPerms := PermRead | PermWrite | PermExecute | PermAdmin
	fmt.Println("admin has write?", adminPerms.Has(PermWrite))

	// ----- 7. role -----
	fmt.Println("guest can edit?", RoleGuest.CanEdit())
	fmt.Println("admin can edit?", RoleAdmin.CanEdit())
	fmt.Println("role name:", RoleModerator)

	// ----- 8. payment method -----
	fmt.Println("esewa online?", PayEsewa.IsOnline())
	fmt.Println("cash online?", PayCash.IsOnline())

	// ----- 9. priority / log -----
	fmt.Println("ticket priority:", PriorityHigh)
	fmt.Println("high > medium?", PriorityHigh > PriorityMedium)
	fmt.Println("log level error:", LogError)

	// ----- 10. environment -----
	appEnv := EnvDev
	fmt.Println("env:", appEnv, "prod?", appEnv.IsProd())

	// ----- 11. weekday -----
	fmt.Println("friday weekend?", Friday.IsWeekend())
	fmt.Println("sunday weekend?", Sunday.IsWeekend())

	// ----- 12. direction + ticket state machine -----
	d := DirNorth
	fmt.Println("start", d, "turn right →", d.TurnRight())

	fmt.Println("open → in progress?", TicketOpen.CanGo(TicketInProgress))
	fmt.Println("open → resolved?", TicketOpen.CanGo(TicketResolved))
	fmt.Println("closed → open?", TicketClosed.CanGo(TicketOpen))

	// ----- 13. tri-state (not bool) -----
	answer := TriUnknown
	fmt.Println("newsletter opt-in:", answer)

	// ----- 14. slice of enums -----
	pipeline := []OrderStatus{Received, Confirmed, Prepared, Delivered}
	fmt.Println("order pipeline:")
	for _, step := range pipeline {
		fmt.Println(" ", step)
	}
}
