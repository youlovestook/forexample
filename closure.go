package main

import "fmt"

// closure is limiting the scope of a variable
// limiting scope is very important

var juicy int // package level scope

func main() {
	foo()
	fmt.Println(juicy)
	{

		yyz := 42
		fmt.Println(yyz)
	}

	fmt.Println(yyz)

}

func foo() {
	fmt.Println("hello")
	juicy++
}
