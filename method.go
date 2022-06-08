package main

import "fmt"

// Class
type pc struct {
	brand  string
	memory int
}

// Method
func (instancePC pc) showMemory() {
	fmt.Println(instancePC.memory)
}

func main() {
	// Instance
	myPC := pc{brand: "Kingston", memory: 2}
	myPC.showMemory()
}
