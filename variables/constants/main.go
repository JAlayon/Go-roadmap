package main

import "fmt"

func main() {
	const days = 7
	const pi float64 = 3.14159
	const secondsInHour = 3600

	duration := 234 // in hours
	fmt.Printf("Duration in seconds %v\n", secondsInHour*duration)

	const n, m int = 4, 5
	const n1, m1 = 6, 7

	const (
		min1 = -500
		min2 = -300
		min3 = 100
	)
	fmt.Printf("Min1: %d, Min2: %d, Min3: %d\n", min1, min2, min3)

}
