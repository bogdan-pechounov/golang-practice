package main

import (
	"fmt"
)

func t1() {
	// grades := [3]int{90, 80, 85}
	grades := [...]int{90, 80, 85}
	var students [3]string = [3]string{"Lisa"}
	students[1] = "John"
	fmt.Println(grades)
	fmt.Println(students)
	fmt.Println(students[2])
}

func t2() {
	a := [...]int{1, 2, 3}
	b := a
	c := &a
	a[1] = 5
	b[0] = 10

	fmt.Println(a, len(a))
	fmt.Println(b)
	fmt.Println(c)
}

func t3() {
	//slices are reference types
	a := []int{1, 2, 3}
	b := a
	// c := &a
	a[1] = 5
	b[0] = 10

	fmt.Println(a, len(a), cap(a))
	fmt.Println(b)
}

func t4() {
	a := []int{1, 2, 3, 4, 5}
	b := a[:3]
	b[0] = 10

	fmt.Println(a, len(a), cap(a))
	fmt.Println(b, len(b), cap(b))
}

func t5() {
	a := []int{}
	fmt.Println(a, len(a), cap(a))
	a = append(a, 0)
	fmt.Println(a, len(a), cap(a))
	a = append(a, 1, 2, 3, 4, 5, 6)
	fmt.Println(a, len(a), cap(a))
}

func t6() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	b := append(a[:2], a[4:]...)

	fmt.Println(b)
	fmt.Println(a)
}

func t7() {
	a := make([]int, 10)
	b := make([]int, 10, 100)

	fmt.Println(len(a), cap(a))
	fmt.Println(len(b), cap(b))
}

func main() {
	t7()
}
