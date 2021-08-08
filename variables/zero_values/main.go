package main

import "fmt"

func main() {
	var intValue = 2
	var floatValue = 5.2

	//Casting
	intValue = int(floatValue)
	fmt.Println(intValue, floatValue)

	var value int
	var price float64
	var name string
	var idDone bool

	fmt.Println(value, price, name, idDone)
}
