package main

import (
	"fmt"
	"math"
)

type figures interface {
	area() float64
}

type circle struct {
	radio float64
}

type rectangle struct {
	base   float64
	height float64
}

func (objectCircle circle) area() float64 {
	return math.Pow(objectCircle.radio, 2) * math.Pi
}

func (objectRectangle rectangle) area() float64 {
	return objectRectangle.base * objectRectangle.height
}

func main() {
	myCircle := circle{radio: 10}
	myRectangle := rectangle{base: 10, height: 18.12}
	fmt.Println(myCircle.area())
	fmt.Println(myRectangle.area())
}
