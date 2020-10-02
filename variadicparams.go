package main

import "fmt"

func main() {

	x := foo(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("the total is : ", x)
}

func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position,", i, "we are adding ", v, " to the total ", sum)
	}

	return sum
}

// func (r reciever) identifier (parameter(s)) (return(s)) { ... all the code }
// stores variadic parameters as a slice
