package shared_lib

import "fmt"

func PrintMessage(msg string) {
	fmt.Println("Shared Library:", msg)
}

// add function to add two numbers
func Add(a, b int) int {
	return a + b
}

// add function to subtract two numbers
func Subtract(a, b int) int {
	return a - b
}
