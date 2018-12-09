package main

import "fmt"

func main() {
	var a = 4
	var b = 5
	var ptr *int

	fmt.Println("The result of a + b is", a + b)

	a += b
	fmt.Println("The new value of a is", a)

	ptr = &a
	fmt.Println("The pointer stores the address", ptr, "which stores the value", *ptr)
}
