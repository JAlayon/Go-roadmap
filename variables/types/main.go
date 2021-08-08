package main

import "fmt"

func main() {
	//INT TYPE
	var i1 int8 = 100
	fmt.Printf("%T\n", i1)

	var i2 uint16 = 65535
	fmt.Printf("%T\n", i2)

	//FLOAT TYPE
	var f1, f2, f3 float64 = 1.1, -0.2, 5.
	fmt.Printf("%T %T %T\n", f1, f2, f3)

	//RUNE TYPE
	var r rune = 'f'
	fmt.Printf("%T\n", r)
	fmt.Println(r)
	fmt.Printf("%x\n", r)

	//BOOL TYPE
	var b bool = true
	fmt.Printf("%T\n", b)

	//STRING TYPE
	var s string = "Hello Go!"
	fmt.Printf("%T\n", s)

	//ARRAY TYPE
	var numbers = [4]int{4, 5, -9, 100}
	fmt.Printf("%T\n", numbers)

	//SLICE TYPE
	var cities = []string{"London", "Tokyo", "NY"}
	fmt.Printf("%T\n", cities)

	//MAP TYPE
	balances := map[string]float64{
		"USD": 2332.4,
		"EUR": 511.2,
	}
	fmt.Printf("%T\n", balances)

	//STRUCT TYPE
	type person struct {
		name string
		age  int
	}

	var person1 person
	fmt.Printf("%T\n", person1)

	//POINTER TYPE
	var x int = 2
	ptr := &x //Address of x
	fmt.Printf("%T\n", ptr)

	//FUNCTION TYPE
	fmt.Printf("%T\n", f)
}

func f() {

}
