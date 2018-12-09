package main

import "fmt"

var (
	some int
	thing string
)

// go-lang has type inference, similar to Swift. It also assigns a default value
// to all variables declared without an initial value.
func main() {
	var a int = 10
	var b = 9
	c := 8

	// Swap the two values.
	a, b = b,a

	// Notice: if we declare a variable without its intial value, we have to assign
	// a type to it (and thus it will be automatically initialized to the default
	// value of that type).
	var d bool

	fmt.Println("The sum is", (a + b + c))
	fmt.Println("The default value of d is", d)

	// We can declare multiple variables in one line.
	var e, f = "45", 67

	fmt.Println("The values are", e, "and", f)
	fmt.Println("The address of e is", &e)
}
