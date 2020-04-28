package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("head first to go"))
	fmt.Println("Hello, playground")

	//string literal
	fmt.Println("Hello, \nGo!")
	fmt.Println("Hello, \tGo!")
	fmt.Println("Quotes: \"\"")
	fmt.Println("Backslash: \\")

	//rune literal
	fmt.Println('A')
	fmt.Println('\t')
	fmt.Println('\n')
	fmt.Println('n')

	//boolean
	fmt.Println(true)
	fmt.Println(false)

	//number
	fmt.Println(10)
	fmt.Println(3.141592)

	//logic
	fmt.Println(10/21, 4.5-2.6)
	fmt.Println(4 < 6)
	fmt.Println(2+2 != 5)

	//type
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf('b')) //rune: int32
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Hello, Go!"))
}
