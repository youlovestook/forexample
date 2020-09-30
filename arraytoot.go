package main

// arrays are a building block for slices
// GO spec says to use slices instead of arrays

import "fmt"

func main() {

	var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))
}
