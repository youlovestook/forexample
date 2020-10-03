package main

import "fmt"

func main() {

	defer foo()

	bar()

}

func foo() {
	fmt.Println("i am foo")
}

func bar() {
	fmt.Println("i am bar")
}
