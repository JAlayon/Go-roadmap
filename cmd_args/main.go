package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*
		fmt.Println("os.Args", os.Args)
		fmt.Println("Path: ", os.Args[0])
		fmt.Println("1st argument: ", os.Args[1])
		fmt.Println("2nd argument: ", os.Args[2])
		fmt.Println("Number of items: ", len(os.Args))
	*/
	var result, err = strconv.ParseFloat(os.Args[1], 64)
	fmt.Printf("%T\n", os.Args[1])
	fmt.Printf("%T\n", result)
	_ = err
}
