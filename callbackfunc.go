package main

import "fmt"

func main() {
	ii := []int{3, 4, 5, 6, 7}
	s := sum(ii...)
	fmt.Println("all numbers", s)

	s2 := even(sum, ii...)

	fmt.Println("even numbers: ", s2)

	s3 := odd(sum, ii...)

	fmt.Println("odd numbers: ", s3)

}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int

	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)

}

func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int

	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)

}

// callback is passing a function as an argument
// functions are first class citizens
// referred to as functional programming
// not always something people should do
// difficult to read
