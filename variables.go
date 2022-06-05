package main

import "fmt"

func main() {
	fmt.Println()
	//Constants
	const pi = 3.14
	fmt.Println(pi)
	//Variable
	var base int = 10
	fmt.Println(base)
	pow := base * base
	fmt.Println(pow)
	//Zero Values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)
	//Zero Value String
	fmt.Println(c)
}
