package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int = 2
	fmt.Println(reflect.TypeOf(myInt))
	fmt.Println(reflect.TypeOf(float64(myInt)))

	//conversion
	var length float64 = 1.2
	var width int = 2
	fmt.Println("Area is", length*float64(width))
	fmt.Println("length > width", length > float64(width))

	//assign
	var length2 float64 = 5.5
	var width2 int = 6
	length2 = float64(width2)
	fmt.Println(length2)

	//remove numbers below decimal point
	var length3 float64 = 6.78
	var width3 int = 4
	width3 = int(length3)
	fmt.Println(width3)

}
