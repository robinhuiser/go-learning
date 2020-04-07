package main

import "fmt"

func main() {
	sliceOfInts := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, myint := range sliceOfInts {
		if myint%2 == 0 {
			fmt.Printf("%d is even\n", myint)
		} else {
			fmt.Printf("%d is odd\n", myint)
		}
	}
}
