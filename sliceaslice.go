package main

import "fmt"

// slice a slice by using the colon operator

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[1:3]) // start at one slice all up to position 3

	for i := 0; i < len(x); i++ { // dont iterate over slices like this use range!
		fmt.Println(x[i])
	}
}
