package main

import "fmt"

var firstint int

type hamburger int

var secondint hamburger

func main() {

	fmt.Println(secondint)
	fmt.Printf("%T\n", secondint)
	secondint = 42
	fmt.Println("second time after assigning   ", secondint)
}
