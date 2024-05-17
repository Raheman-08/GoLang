package main

import "fmt"

func main() {
	fmt.Println("Welcome to the class of pointers")

	// var ptr *int

	// fmt.Println("value of pointer is ", ptr)

	myNumber := 40

	var ptr = &myNumber

	fmt.Println("The actual pointer is ", *ptr)
	fmt.Println("The actual pointer is ", ptr)

	*ptr = *ptr * 2

	fmt.Println("New value is ", myNumber)
}
