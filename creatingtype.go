package main

import "fmt"

// you can create your own type in GO
var aa int

type hotdog int

var bb hotdog

func main() {
	aa = 42
	bb = 43
	fmt.Println(aa)
	fmt.Printf("%T\n", aa)
	fmt.Println(bb)
	fmt.Printf("%T\n", bb)

	// this is called conversion and not casting

	aa = int(bb)
	fmt.Println(aa)
	fmt.Printf("%T\n", aa)
}
