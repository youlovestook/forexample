package main

import "fmt"

var yy = 42

// static programming languange not dynamic
// a variable is declared to hold value of a certain type
// DECLARED that the VARIABLE with the IDENTIFIER "z"
// is of TYPE int

var zz = "Shaken, not stirred"
var a = `James said "shaken not stirred"`

func main() {
	fmt.Println("y")
	fmt.Printf("%T\n", yy)
	fmt.Println("z")
	fmt.Printf("%T\n", zz)
	zz = "43"
	fmt.Println(zz)
	fmt.Printf("%T\n", zz)
	fmt.Printf("%T\n", a)

}
