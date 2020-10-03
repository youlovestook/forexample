package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)
	x := bar()

	fmt.Printf("%T\n", x)
	i := x()
	fmt.Println(i)
}

func foo() string {
	s := "hello world"
	return s
}

func bar() func() int {
	return func() int {
		return 451
	}
}
