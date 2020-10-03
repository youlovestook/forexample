package main

import "fmt"

func main() {
	fmt.Println("eee")
	s := foo()
	fmt.Println(s)
	fmt.Println("-------------")
	j, k := bar()
	fmt.Println(j, "   ", k)
}

func foo() int {
	return 24
}

func bar() (string, int) {
	q := "boogie oogie oogie"
	r := 42
	return q, r
}
