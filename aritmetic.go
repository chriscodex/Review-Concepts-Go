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
	// Subtraction
	result = y - x
	fmt.Println("Subtraction: ", result)
	// Multiplication
	result = x * y
	fmt.Println("Multiplication: ", result)
	// Division
	result = y / x
	fmt.Println("Division: ", result)
	// Module
	result = y % x
	fmt.Println("Module: ", result)
	// Incremental
	x++
	fmt.Println(x)
	// Decremental
	x--
	fmt.Println(x)
	// Circle area
	var r uint = 10
	areaC := math.Pow(float64(r), 2) * math.Pi
	fmt.Println("Area: ", areaC)
	// Trapezium area
	a := 10
	b := 12
	areaT := (a + b) / 2
	fmt.Println(areaT)
}
