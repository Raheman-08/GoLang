package main

import (
	"fmt"
)

func main() {
	var username string = "ali"
	fmt.Println(username)
	fmt.Printf("Variable is type of: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is type of: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is type of: %T \n", smallVal)
}
