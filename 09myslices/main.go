package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to the slice in golang")

	var fruitList = []string{"Apple", "Banana", "Peach"}
	fmt.Printf("Type of fruitlist %T\n", fruitList)

	fruitList = append(fruitList, "Kiwi", "Mango")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 123
	highScore[1] = 345
	highScore[2] = 543
	highScore[3] = 233
	// highScore[4] = 654

	highScore = append(highScore, 666, 556, 765)

	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))
}
