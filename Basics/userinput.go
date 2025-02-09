package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Asking for user name
	fmt.Print("What's your name? ")

	// Using bufio for safe input handling
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	// Printing greeting message
	fmt.Println("Hello,", name)
}
