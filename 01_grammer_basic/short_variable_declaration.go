package main

import "fmt"

func main() {
	quantity := 4
	length, width := 3.6, 5.9
	customerName := "Jungyeon Lee"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")
}
