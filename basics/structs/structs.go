package main

import (
	"fmt"
)

type Animal struct {
	Name   string
	Origin string
	test   string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
	test2    string
}

func main() {
	bird := Bird{}
	fmt.Println(bird)

	bird.Name = "Emu"
	bird.test = "TEST"
	bird.test2 = "TEST2"

	fmt.Println(bird)

}
