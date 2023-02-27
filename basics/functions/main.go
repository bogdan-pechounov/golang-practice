package main

import (
	"fmt"
)

func sayMessage(greeting, name string) {
	fmt.Println(greeting, name)
	name = "Ted"
	fmt.Println(name)
}

func sum(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

func t() {
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

func t2() {
	func() {
		fmt.Println("Hello Go!")
	}()
}

func t3() {
	var divide func(float64, float64) (float64, error)
	divide = func(a, b float64) (float64, error) {
		if b == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return a / b, nil
	}
	d, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

func t4() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
}

func main() {
	t4()
}
