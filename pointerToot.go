package main

import "fmt"

func main() {

	a := 42

	fmt.Println(a)         // value
	fmt.Println(&a)        // aaaaand what is the address
	fmt.Printf("%T\n", a)  // value
	fmt.Printf("%T\n", &a) // we are pointing to an int

	b := &a
	fmt.Println(b)
	fmt.Println(*b) // gives us the value stored at that address
	fmt.Println(*&a)
}

// everything in go is passed by value
