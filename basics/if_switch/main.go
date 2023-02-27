package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("true")
	}

	if a := 5; a > 6 {
		fmt.Println(a)
	} else if b := 4; a > b {
		fmt.Println(a, b)
	}

	switch 2 {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("default")
	}

	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to 10")
	}
}
