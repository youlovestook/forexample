package main

import "fmt"

// slice is dynamic
// we made a slice of 10
// the underlying array is of size 100

func main() {
	x := make([]int, 10, 100)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 1
	x[9] = 77
	fmt.Println(x)

	x = append(x, 3423)
	fmt.Println(x)
	fmt.Println(len(x))
}
