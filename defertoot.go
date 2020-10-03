package main

import "fmt"

func main() {

	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

// defer is the last thing to happen before main closes
