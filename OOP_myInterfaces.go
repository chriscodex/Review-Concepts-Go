package main

import (
	"fmt"
	inter "reviewConceptsGo/myInterfaces"
)

func main() {
	myPeripheral := inter.Peripheral{Brand: "Asus", Model: "K555U"}
	myDog := inter.Dog{Name: "Rocky", Size: "Medium"}
	fmt.Println(myPeripheral.ShowFirstItem())
	fmt.Println(myDog.ShowFirstItem())
	/* Output in console:
	Asus
	Rocky
	*/
	// Both objects have the same method wich is ShowFirstItem, but with interfaces the program can difference them
	
}
