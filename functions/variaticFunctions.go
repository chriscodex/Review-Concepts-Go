package main

import "fmt"

func addit(values ...int) int {
	total := 0
	for _, element := range values {
		total += element
	}
	return total
}

func main() {
	fmt.Println(addit(1, 2))
	fmt.Println(addit(1, 2, 3, 4))
}
