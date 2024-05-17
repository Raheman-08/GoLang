package main

import "fmt"

func main() {
	fmt.Println("This is an array class")

	var fruitList [7]string

	fruitList[0] = "apple"
	fruitList[2] = "banana"
	fruitList[6] = "watermelon"

	fmt.Println("Fruit list is ", len(fruitList))
}
