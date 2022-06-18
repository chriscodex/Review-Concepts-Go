package main

import "fmt"

func simpleExample() {
	x := "hello"
	y := func() string {
		return x + " world"
	}()
	fmt.Println(y)
}

func main() {
	simpleExample()
}
