package main

import "fmt"

func main() {
	numSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	oddOrEven(numSlice)
}

func oddOrEven(n []int) {
	for _, num := range n {
		if num%2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}
}
