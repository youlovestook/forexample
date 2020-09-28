package main

import "fmt"

var y = 43

var z int
var zztop string

func main() {
	//short declaration operator
	// DECLARE a variable and ASSIGN a VALUE (of a certain type)
	// var keyword changes the scope of the variable

	x := 42
	fmt.Println(x)

	fmt.Println(y)

	footie()

	fmt.Println(z)
	fmt.Println(zztop)
}

func footie() {
	fmt.Println("again:", y)
}
