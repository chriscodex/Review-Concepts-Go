package main

import (
	"fmt"
	"time"
)

func simpleExample() {
	x := "hello"
	y := func() string {
		return x + " world"
	}()
	fmt.Println(y)
}

func goRoutineExample() {
	c := make(chan struct{})
	go func() {
		fmt.Println("Starting")
		time.Sleep(3 * time.Second)
		fmt.Println("End")
		c <- struct{}{}
	}()
	<-c
}

func main() {
	simpleExample()
	goRoutineExample()
}
