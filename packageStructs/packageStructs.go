package packageStructs

import "fmt"

type PC struct {
	Brand string
	Ram   int
}

func (objectPc PC) ShowMemory() {
	fmt.Println(objectPc.Ram)
}

func (objectPc *PC) DuplicateRam() {
	objectPc.Ram = objectPc.Ram * 2
	fmt.Println("Memory Ram has been duplicated")
}
