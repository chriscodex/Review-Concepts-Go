package main

import (
	"fmt"
	"math"
)

func main() {
	x := 5
	y := 15
	// Addition
	result := x + y
	fmt.Println("Addition: ", result)
	/*Output
	20
	*/
	// Subtraction
	result = y - x
	fmt.Println("Subtraction: ", result)
	/*Output
	10
	*/
	// Multiplication
	result = x * y
	fmt.Println("Multiplication: ", result)
	/*Output
	75
	*/
	// Division
	result = y / x
	fmt.Println("Division: ", result)
	/*Output
	3
	*/
	// Module
	result = y % x
	fmt.Println("Module: ", result)
	/*Output
	0
	*/
	// Incremental
	x++
	fmt.Println(x)
	/*Output
	6
	*/
	// Decremental
	x--
	fmt.Println(x)
	/*Output
	5
	*/
	// Circle area
	var r uint = 10
	areaC := math.Pow(float64(r), 2) * math.Pi
	fmt.Println("Area: ", areaC)
	/*Output
	Area:  314.1592653589793
	*/
	// Trapezium area
	a := 10
	b := 12
	areaT := (a + b) / 2
	fmt.Println(areaT)
	/*Output
	11
	*/
}
