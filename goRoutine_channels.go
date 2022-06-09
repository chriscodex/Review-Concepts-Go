package main

import "fmt"

func say(text string, c chan<- string) {
	// Into the channel
	c <- text
}

func main() {
	cs := make(chan string, 1)

	fmt.Println("Hello")

	go say("Bye", cs)

	// Out of the channel
	fmt.Println(<-cs)
}
