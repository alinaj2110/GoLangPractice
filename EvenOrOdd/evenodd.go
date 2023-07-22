package main

import "fmt"

func main() {
	numberSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, n := range numberSlice {
		if n%2 == 0 {
			fmt.Println(n, " - Even")
		} else {
			fmt.Println(n, " - Odd")
		}
	}
}
