package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is the price of Coffee:")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for pricing: ", input)
	fmt.Printf("The type of pricing %T ", input)
}
