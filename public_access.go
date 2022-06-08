package main

import (
	"fmt"
	"reviewConceptsGo/myPackage"
	pk "reviewConceptsGo/myPackage"
)

func main() {
	var personPublic myPackage.PersonPublic
	personPublic.Name = "Christian"
	personPublic.Age = 24
	fmt.Println(personPublic)
	pk.PrintMessage("This is a message")
}
