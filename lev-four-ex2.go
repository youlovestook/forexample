package main

import "fmt"

func main() {

	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
