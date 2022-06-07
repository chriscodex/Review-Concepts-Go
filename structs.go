package main

import "fmt"

type person struct {
	name      string
	last_name string
	age       int
}

func main() {
	person1 := person{name: "Christian", last_name: "Espinoza", age: 23}
	fmt.Println(person1)
	// Other way
	var person2 person
	person2.name = "Leo"
	fmt.Println(person2)
}
