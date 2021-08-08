package main

import (
	"fmt"
)

func main() {
	const (
		c1 = iota
		c2 = iota
		c3 = iota
		c4 = iota
	)

	fmt.Println(c1, c2, c3, c4)

	const (
		a1 = iota
		a2
		a3
	)

	fmt.Println(a1, a2, a3)

	const (
		North = iota //default 0
		East
		South
		West
	)
	fmt.Println(West) // 3

	const (
		a = iota * 2 // 0
		b            // 2
		c            // 4
	)
	fmt.Println(a, b, c)

	// How to skip some values
	const (
		x = -(iota + 2) //-2
		_               //-3
		y               //-4
		z               //-5
	)
	fmt.Println(x, y, z)
}
