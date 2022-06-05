package main

import (
	"fmt"
	"log"
	"strconv"
)

func pairOrOdd(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	value1, value2 := 3, 7
	// And
	if value1 == 3 && value2 == 5 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	// Or
	if value1 == 3 || value2 == 5 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	// Parse string to int
	value, err := strconv.Atoi("33")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(value)
	}
	// Pair or Odd
	number := 7
	result := pairOrOdd(number)
	if result {
		fmt.Printf("The number %d is pair", number)
	} else {
		fmt.Printf("The number %d is odd\n", number)
	}
	// Normal Switch
	val3 := 100
	switch val3 {
	case 10:
		fmt.Println("Ten")
	case 100:
		fmt.Println("One Hundred")
	default:
		fmt.Println("Other number")
	}
	// Other way (Only for switch scope)
	switch val4 := 10; val4 {
	case 10:
		fmt.Println("Ten")
	case 100:
		fmt.Println("One Hundred")
	default:
		fmt.Println("Other number")
	}
	// No condition
	val4 := 6
	switch {
	case val4 < 5:
		fmt.Println("Less than 5")
	case val4 > 10:
		fmt.Println("Higher than 10")
	default:
		fmt.Println("Other number")
	}
}
