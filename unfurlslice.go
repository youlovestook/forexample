package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x := foo(xi...) // will unfurl the ints and pass them in since go is static
	fmt.Println("the total is : ", x)
}

func foo(x ...int) int { // variadic means you can pass 0 to an unlimited number of values
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position,", i, "we are adding ", v, " to the total ", sum)
	}

	return sum
}

// func (r reciever) identifier (parameter(s)) (return(s)) { ... all the code }
// stores variadic parameters as a slice
