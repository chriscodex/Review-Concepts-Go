package main

import "fmt"

func main() {
	// Arrays
	array := [6]int{0, 1, 2, 3, 4, 5}
	fmt.Println(array)
	fmt.Println(len(array), cap(array))
	// Slice
	slice := []int{0, 1, 2, 3}
	fmt.Println(slice)
	slice2 := slice[:2]
	fmt.Println(slice2, len(slice2))
	slice = append(slice, 4)
	fmt.Println(slice)
	slice3 := []int{5, 6}
	slice = append(slice, slice3...)
	fmt.Println(slice)
	slice[0] = 10
	fmt.Println(slice, slice2)
	for i := range array {
		fmt.Println(i)
	}
}
