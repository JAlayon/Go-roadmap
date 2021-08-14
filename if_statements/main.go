package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	num, err := strconv.Atoi("abc")
	_ = err

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Num: ", num)
	}

	if num, err := strconv.Atoi("20"); err == nil {
		fmt.Println("No error, num is: ", num)
	} else {
		fmt.Println(err)
	}

	if args := os.Args; len(args) != 2 {
		fmt.Println("One argument is required")
	} else if km, err := strconv.ParseFloat(args[1], 64); err == nil {
		fmt.Printf("%v km in miles is %v\n", km, km/1.609)
	}

}
