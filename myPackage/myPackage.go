package myPackage

import "fmt"

// PersonPublic Person with public access
type PersonPublic struct {
	Name string
	Age  int
}

// Public Functions
func PrintMessage(text string) {
	fmt.Println(text)
}
