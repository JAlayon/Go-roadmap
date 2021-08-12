package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := string(99) //c
	fmt.Println(s)

	var myStr = fmt.Sprintf("%f", 44.2)
	fmt.Println(myStr)

	var myStr1 = fmt.Sprintf("%d", 19)
	fmt.Println(myStr1)

	s1 := "3.123"
	fmt.Printf("%T\n", s1)

	var f1, err = strconv.ParseFloat(s1, 64)
	_ = err
	fmt.Println(f1)

	i, err := strconv.Atoi("50")
	s2 := strconv.Itoa(i)
	fmt.Printf("i type is %T, i value is %v\n", i, i)     // int, 50
	fmt.Printf("s2 type is %T, s2 value is %v\n", s2, s2) //string 50

}
