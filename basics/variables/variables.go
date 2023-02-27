package main

import (
	"fmt"
)

func test1() {
	a := 10
	var b int = 3
	c := a + b
	fmt.Println(c)
	fmt.Println(a & b)
	// var d uint = -1
}

func test2() {
	var n complex64 = 1 + 2i
	r := real(n)
	i := imag(n)
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", r, r)
	fmt.Printf("%v, %T\n", i, i)

	a := 1 + 2i
	b := 3 + 4i
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Print(a * b)
	fmt.Print(a / b)
}

func test3() {
	s := "this is a string"
	// s[1] = 'a'
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s[3], s[3])
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)
}

// type of constant
func test4() {
	const a = 1
	// const a2 int32 = 2
	const a2 = 2
	var b int16 = 27
	var c uint64 = 1e10

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", a2, a2)

	fmt.Println(a2 + b) //implicit conversion
	fmt.Println(a2 + c)
}

func test5() {
	//iota
	const (
		_ = 1 << iota
		a
		b
		c
	)

	fmt.Println(a, b, c)
}

func main() {
	test5()
}
