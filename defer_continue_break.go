package main

import "fmt"

func deferFunc() {
	// Defer
	fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")
	defer fmt.Println("4")
}

func main() {
	// Defer
	deferFunc()

	// Continue / Break
	for i := 0; i < 10; i++ {
		if i == 2 {
			println(i)
			continue
		}
		if i == 6 {
			fmt.Println(i)
			break
		}
	}
}
