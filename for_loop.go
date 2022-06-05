package main

import "fmt"

func main() {
	// For conditional
	fmt.Println("For conditional")
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	// For while
	fmt.Println("For while")
	counter := 0
	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}
	// For range
	fmt.Println("For range")
	pairNumbers := []int{2, 4, 6}
	for i, element := range pairNumbers {
		fmt.Printf("Position %d: %d\n", i, element)
	}

	// For forever
	/*for {
		fmt.Println(counter)
		counter++
	}*/

}
