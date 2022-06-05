package main

import "fmt"

// No return
func printMessage(mess string, num int) {
	fmt.Printf("%s has received %d messages", mess, num)
}

// Return
func returnValue(a int) int {
	return a * a
}

//Double Return
func doubleReturn(a int) (int, string) {
	b := a * 3
	return b, "messages"
}
func doubleReturn2(a int) (int, int, string) {
	return a * 8, a * 7, "hello"
}

func main() {
	// No return function
	printMessage("Assistant Virtual", 100)
	// Return function
	pow := returnValue(5)
	fmt.Printf("\nFunction returnValue: %d", pow)
	// Double return function
	value1, value2 := doubleReturn(4)
	fmt.Printf("\nValue1: %d\nValue2: %s", value1, value2)
	// Double return 2 function
	value3, value4, value5 := doubleReturn2(1)
	fmt.Printf("\nValue1: %d\nValue2: %d\nValue3: %s", value3, value4, value5)
}
