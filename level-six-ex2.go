package main

import "fmt"

func main() {
	x := []int{2, 4, 6, 8, 10}
	n := foo(x...)
	fmt.Println("the total is : ", n)

	fmt.Println("-------------bar------------")
	xr := bar([]int{2, 4, 6, 8, 10})
	fmt.Println("the total of bar is : ", xr)
}

func foo(x ...int) int {

	sum := 0

	for _, v := range x {
		sum += v
		//fmt.Println("for item in index position,", i, "we are adding ", v, " to the total ", sum)
	}

	return sum
}

func bar(x []int) int {

	sum := 0

	for _, v := range x {
		sum += v
		//fmt.Println("for item in index position,", i, "we are adding ", v, " to the total ", sum)
	}

	return sum
}

// func (r reciever) identifier (parameter(s)) (return(s)) { ... all the code }
// stores variadic parameters as a slice
