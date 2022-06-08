package main

import (
	"fmt"
	"sync"
	"time"
)

func prtMessage(txt string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(txt)
}

func main() {
	var wg sync.WaitGroup

	start := time.Now()
	fmt.Println("Hello")
	wg.Add(1)
	go prtMessage("World", &wg)
	wg.Wait()
	fmt.Println(time.Since(start))
}
