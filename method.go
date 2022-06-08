package main

import (
	"fmt"
)

// Class
type pc struct {
	brand  string
	memory int
}

// Method
func (instancePC pc) showMemory() {
	fmt.Println(instancePC.memory)
}

func (instancePC2 *pc) duplicateRam() {
	instancePC2.memory = instancePC2.memory * 2
	fmt.Println(instancePC2.memory)
}

func main() {
	// Object
	myPC := pc{brand: "Kingston", memory: 2}
	myPC.showMemory()
	myPC.duplicateRam()
	fmt.Println(myPC.memory)
}
