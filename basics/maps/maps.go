package main

import (
	"fmt"
)

func t1() {
	statePopulations := map[string]int{
		"Florida":  206,
		"New York": 197,
	}
	statePopulations["Georgia"] = 103
	fmt.Println(statePopulations, statePopulations["Georgia"])
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations, statePopulations["Georgia"])

	_, ok := statePopulations["Ohio"]
	fmt.Println(ok)
}

type Doctor struct {
	number     int
	name       string
	companions []string
}

func t2() {
	aDoctor := Doctor{
		number:     3,
		name:       "John",
		companions: []string{"Liz", "Jo"},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.name)
}

func main() {
	t2()
}
