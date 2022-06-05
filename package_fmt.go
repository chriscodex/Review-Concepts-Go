package main

import "fmt"

func main() {
	// Variable
	chain := "Golang"
	integ := 10
	// Line break
	fmt.Println("Linea break", chain, integ)
	// Template String
	fmt.Printf("String: %s\nInt: %d", chain, integ)
	// Template String for variables
	fmt.Printf("\nVariable: %v", chain)
	// Sprintf
	message := fmt.Sprintf("New message: %d", integ)
	fmt.Println()
	fmt.Println(message)
	// Datatype
	fmt.Printf("Datatype of integ variable: %T", integ)
}
