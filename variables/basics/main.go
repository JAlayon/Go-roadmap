package main

import "fmt"

func main() {

	//** Declaring Variables**//

	//Syntax: var var_name type
	var s1 string
	s1 = "Learning go!"
	fmt.Println(s1)

	//** Type Inference **//
	var age int = 30 //Not neccessary to say the type (int)
	var i = 10
	var j = 20
	fmt.Println("Age:", age)
	fmt.Println(i, j)

	var name = "Dan"
	//fmt.Println("Your name is:", name)
	_ = name // In Go each varible must be used or will we have a compile-time error
	//_ is the Blank Identifier and mutes the error of unused variables

	//** Short Declaration (works only in Block Scope) **//
	s := "Learning golang"
	fmt.Println(s)

	//Multiple Declarations
	car, cost := "Audi", 10000
	fmt.Println(car, cost)
	car, year := "BMW", 2021
	fmt.Println(car, year)

	var opened = false
	opened, file := true, "a.txt"
	_, _ = opened, file

	var (
		salary    float64
		firstName string
		gender    bool
		//_,_,_ = salary,firstName,gender
	)

	fmt.Println(salary, firstName, gender)

	var a, b, c, d int
	_, _, _, _ = a, b, c, d

	//Multiple Assignments
	var ii, jj int
	ii, jj = 5, 8
	_, _ = ii, jj

	jj, ii = ii, jj //swapping variables
	fmt.Println(ii, jj)

	sum := 5 + 2.3
	fmt.Println("Sum:", sum)

}
