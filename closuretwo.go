package main

import "fmt"

// closure is limiting the scope of a variable
// limiting scope is very important

var juicy int // package level scope

func main() {
	a := incrementor()
	fmt.Println(a())

}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x

	}
}
