package main

import "fmt"

type creature struct {
	name string
	sort string
}

func (objectCriature creature) String() string {
	return fmt.Sprintf("The name of the criature is %s and its sort is %s", objectCriature.name, objectCriature.sort)
}

func main() {
	crearture1 := creature{name: "Wet", sort: "Water"}
	fmt.Println(crearture1)
}
