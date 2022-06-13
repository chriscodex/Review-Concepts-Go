package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func scan() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Text()
	return scanner.Text()
}

func isPalindrome(word string) bool {
	word = strings.ToLower(word)
	word = strings.ReplaceAll(word, " ", "")
	var textReverse string
	for i := len(word) - 1; i >= 0; i-- {
		textReverse += string(word[i])
	}
	if textReverse == word {
		fmt.Println("It's a palindrome")
	} else {
		fmt.Println("It's not a palindrome")
	}
}

func main() {
	scan()
	if isPalindrome("amor a roma") {
		fmt.Println("It's a palindrome")
	} else {
		fmt.Println("It's not a palindrome")
	}
}
