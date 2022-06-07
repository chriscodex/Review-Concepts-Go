package main

import "fmt"

func main() {
	maps := make(map[string]string)
	maps["name"] = "Christian"
	maps["last name"] = "Espinoza"
	fmt.Println(maps)
	for _, value := range maps {
		fmt.Println(value)
	}
	// Find value
	_, ok := maps[""]
	fmt.Println(ok)
	for k, v := range maps {
		fmt.Println(k, v)
	}
}
