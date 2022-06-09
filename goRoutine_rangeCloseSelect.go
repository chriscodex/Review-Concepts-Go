package main

import "fmt"

func msg(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "Message 1"
	c <- "Message 2"
	fmt.Println(len(c), cap(c))
	close(c)
	for i := range c {
		fmt.Println(i)
	}

	// Select
	email1 := make(chan string)
	email2 := make(chan string)
	go msg("message1", email1)
	go msg("message2", email2)
	for i := 0; i < 2; i++ {
		select {
		case <-email1:
			fmt.Println("Email recieved from email1")
		case <-email2:
			fmt.Println("Email recieved from email2")
		}
	}
}
