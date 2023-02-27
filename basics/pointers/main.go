package main

import (
	"fmt"
)

func t1() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, *b)
	a = 27
	fmt.Println(a, *b)
}

type myStruct struct {
	foo int
}

func t2() {
	var ms *myStruct
	fmt.Println(ms)
	ms = new(myStruct)
	(*ms).foo = 42
	fmt.Println(ms)
}

func main() {
	t2()
}
