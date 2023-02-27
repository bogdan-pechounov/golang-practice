package main

import (
	"fmt"
)

func main() {
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
	fmt.Println("after:", i)

	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}
}
