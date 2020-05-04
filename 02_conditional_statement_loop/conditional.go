package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("I'll be printed!")
	}

	var grade int = 100
	if grade == 100 {
		fmt.Println("Perfect!")
	} else if grade >= 60 {
		fmt.Println("You pass.")
	} else {
		fmt.Println("You fail!")
	}

	if 1 == 1 {
		fmt.Println("I'll be printed.")
	}

	if 1 > 2 {
		fmt.Println("I won't.")
	}
	if 1 < 2 {
		fmt.Println("I'll be printed!")
	}

	if !true {
		fmt.Println("I won't be printed!")
	}

	// and
	if true && true {
		fmt.Println("I'll be printed.")
	}
	if true && false {
		fmt.Println("I won't!")
	}

	// or
	if false || true {
		fmt.Println("I'll be printed!")
	}
	if false || false {
		fmt.Println("I won't.")
	}
}
