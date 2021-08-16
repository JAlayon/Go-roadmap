package main

import (
	"fmt"
	"time"
)

func main() {
	language := "golang"

	switch language {
	case "python":
		fmt.Println("You are learning Python")
	case "Go", "golang":
		fmt.Println("You are learning golang")

	default:
		fmt.Println("Any other programming language is a good option")
	}

	n := 5
	switch true {
	case n%2 == 0:
		fmt.Println("Even number!")
	case n%2 != 0:
		fmt.Println("Odd number!")
	default:
		fmt.Println("Never here!")
	}

	var time = time.Now()
	_ = time
	var hour = time.Hour()

	switch {
	case hour < 12:
		fmt.Println("Good Morning")
	case hour < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Godd evening")
	}
}
