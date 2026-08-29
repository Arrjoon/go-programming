package main

import "fmt"

// maps -> hash table / dictionary / object
func main() {
	// 1. Declare only (nil map — cannot add keys yet)
	var nilMap map[string]int
	fmt.Println("nil map:", nilMap)
	fmt.Println("is nil?", nilMap == nil)
	fmt.Println("len of nil map:", len(nilMap))

	// 2. make — empty map you can write to
	m := make(map[string]string)
	m["name"] = "golang"
	m["area"] = "backend"
	fmt.Println("make map:", m)
	fmt.Println("get name:", m["name"])

	// 3. make with a size hint (optional, for performance)
	hinted := make(map[string]int, 10)
	hinted["a"] = 1
	fmt.Println("make with hint:", hinted)

	// 4. Map literal — create and fill in one go
	ages := map[string]int{
		"alice": 24,
		"bob":   30,
	}
	fmt.Println("literal map:", ages)

	// 5. Empty map literal (not nil — you can add keys)
	empty := map[int]string{}
	empty[1] = "one"
	fmt.Println("empty literal:", empty)

	// 6. var + later assignment
	var scores map[string]int
	scores = map[string]int{"math": 90, "science": 85}
	fmt.Println("assigned later:", scores)

	// 7. Get a value — missing key returns the zero value
	fmt.Println("existing key:", ages["alice"])
	fmt.Println("missing key (zero value):", ages["carol"])

	// 8. Check if a key exists (ok idiom)
	value, ok := ages["alice"]
	fmt.Println("alice exists?", ok, "value:", value)

	missing, found := ages["carol"]
	fmt.Println("carol exists?", found, "value:", missing)

	if age, exists := ages["bob"]; exists {
		fmt.Println("bob's age:", age)
	}

	// 9. Update a value
	ages["alice"] = 25
	fmt.Println("after update:", ages)

	// 10. Add a new key
	ages["carol"] = 22
	fmt.Println("after add:", ages)

	// 11. Delete a key
	delete(ages, "bob")
	fmt.Println("after delete bob:", ages)

	// 12. Length
	fmt.Println("len:", len(ages))

	// 13. Loop with range (key and value)
	fmt.Println("range key + value:")
	for name, age := range ages {
		fmt.Println(name, "->", age)
	}

	// 14. Range — keys only
	fmt.Println("range keys only:")
	for name := range ages {
		fmt.Println(name)
	}

	// 15. Range — values only
	fmt.Println("range values only:")
	for _, age := range ages {
		fmt.Println(age)
	}

	// 16. Different key / value types
	idToName := map[int]string{1: "ram", 2: "sita"}
	fmt.Println("int -> string:", idToName)

	flags := map[string]bool{"ready": true, "done": false}
	fmt.Println("string -> bool:", flags)

	// 17. Map of slices
	groups := map[string][]string{
		"frontend": {"html", "css", "js"},
		"backend":  {"go", "sql"},
	}
	fmt.Println("map of slices:", groups)
	fmt.Println("backend[0]:", groups["backend"][0])

	// 18. Nested maps (map of maps)
	users := map[string]map[string]string{
		"u1": {"name": "arjun", "role": "dev"},
		"u2": {"name": "maya", "role": "qa"},
	}
	fmt.Println("nested map:", users)
	fmt.Println("u1 name:", users["u1"]["name"])

	// 19. Struct as a value
	type person struct {
		name string
		age  int
	}
	people := map[string]person{
		"a": {name: "alice", age: 24},
		"b": {name: "bob", age: 30},
	}
	fmt.Println("map of structs:", people)
	fmt.Println("alice age:", people["a"].age)

	// 20. Compare maps only with nil (not with each other)
	var a map[string]int
	b := map[string]int{}
	fmt.Println("nil == nil map?", a == nil)
	fmt.Println("empty literal is nil?", b == nil)

	// 21. Clear all keys (Go 1.21+)
	toClear := map[string]int{"x": 1, "y": 2}
	clear(toClear)
	fmt.Println("after clear:", toClear, "len:", len(toClear))
}
