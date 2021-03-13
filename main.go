package main

import (
	"fmt"
)

func main() {
	listOfInts := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, slot := range listOfInts {
		if slot%2 == 0 {
			fmt.Println("This slot is even:", slot)
		} else {
			fmt.Println("this slot is odd:", slot)
		}
	}

}
