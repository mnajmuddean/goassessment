package main

import (
	// import all the packages
	"fmt"
	"strings"
)

func main() {
	// main function
	fmt.Println("Please enter your name.")
	// declare variable string named name
	var name string
	// tells the computer to wait for input
	// from the keyboard ending with a new line or (\n), character
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)
	fmt.Printf("Hi, %s! I'm Go!", name)
}
