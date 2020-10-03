package main

import "fmt"

func main() {

	foo()

	func(x int) {
		fmt.Println("anonymous ", x)
	}(42)
}

func foo() {
	fmt.Println("foo ran")
}
