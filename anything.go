package main

import "fmt"

func main() {
	fmt.Println("HEllo everyone this is the most awesome class")
	foo()
	fmt.Println("something more")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I'm in Foo")
}

func bar() {
	fmt.Println("we exited")
}

// control flow
// 1. sequential
// 2. loop or iterative
// 3. conditional
