package main

import "fmt"

func main() {
	var a = 100
	var b = 200

	fmt.Printf("The max value is %d.\n", max(a, b))
}

func max(num1 int, num2 int) int {
	if num1 >= num2 {
		return num1
	} else {
		return num2
	}
}
