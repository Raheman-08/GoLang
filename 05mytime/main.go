package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("This is study time of golang")

	presentTime := time.Now()

	fmt.Println(presentTime.Format("01 Jan 2006 Monday"))

	createdDate := time.Date(2000, time.June, 30, 23, 23, 0, 0, time.UTC)

	fmt.Println(createdDate)
}
