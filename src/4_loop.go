package main

import "fmt"

func main() {
	for a := 0; a < 10; a++ {
		fmt.Printf("The current value is %d.\n", a)
	}

	numbers := [6]int{1, 2, 3, 4}
	for i, x := range numbers {
		fmt.Printf("%dth number in the array is %d\n", i, x)
	}
}
