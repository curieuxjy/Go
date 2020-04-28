package main

import (
	"fmt"
	"reflect"
)

func main() {
	var quantity int
	var length, width float64
	var customerName string

	quantity = 4
	length, width = 1.2, 2.4
	customerName = "Jungyeon Lee"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")

	// continued
	// types can be ommitted
	var quantity2 int = 5
	var length2, width2 float64 = 5.8, 6.9
	var customerName2 string = "Curieux JY"
	fmt.Println(reflect.TypeOf(quantity2))
	fmt.Println(reflect.TypeOf(length2))
	fmt.Println(reflect.TypeOf(width2))
	fmt.Println(reflect.TypeOf(customerName2))
}
