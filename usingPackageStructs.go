package main

import (
	ps "reviewConceptsGo/packageStructs"
)

func main() {
	myNewPC := ps.PC{Brand: "MSI", Ram: 2}
	myNewPC.ShowMemory()
	myNewPC.DuplicateRam()
	myNewPC.ShowMemory()
}
