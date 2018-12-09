package main

import "fmt"
import "unsafe"

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

const (
	d = iota
	e
	f = "haha"
	g
	h = 100
	i
	j = iota
	k
	l = 1 << iota
	m
)

func main() {
	const LENGTH int = 10
	// A constant has to be declared with an initial value (since it is constant whose value cannot change anymore).
	const DECISION = false

	fmt.Println("Since the length is", LENGTH, ". The decision is", DECISION)
	fmt.Println(a, b, c)
	fmt.Println(d, e, f, g, h, i, j, k, l, m)
}
