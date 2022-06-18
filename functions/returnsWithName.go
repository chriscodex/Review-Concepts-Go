package main

import "fmt"

func mult(x int) (double int, triple int, quad int) {
	double = x * 2
	triple = x * 3
	quad = x * 4
	return
}

func main() {
	d, t, q := mult(5)
	fmt.Println(d, t, q)
	/*Output
	10 15 20
	*/
}
