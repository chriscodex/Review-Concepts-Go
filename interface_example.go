package main

import "fmt"

type animal interface {
	movement() string
}

type dog struct{}
type bird struct{}
type fish struct{}

func (d dog) movement() string {
	return "Walking"
}
func (b bird) movement() string {
	return "Flying"
}
func (f fish) movement() string {
	return "Swiming"
}
func movementAnimal(a animal) {
	fmt.Println(a.movement())
}

func main() {
	d := dog{}
	b := bird{}
	f := fish{}
	movementAnimal(d)
	movementAnimal(b)
	movementAnimal(f)
}
