package main

import "fmt"

// when to use pointers: when you have a large piece of data that you dont want to pass around
// you can just pass the address
// if you need to change a value at a specific location
// just a reminder everything in go is pass by value

func main() {

	x := 0
	fmt.Println("x before ", &x)
	fmt.Println("x before ", x)
	foo(&x)
	fmt.Println("x after ", &x)
	fmt.Println("x after ", x)
}

func foo(y *int) {

	fmt.Println("y before ", y)
	fmt.Println("y before ", *y)
	*y = 43
	fmt.Println("y after ", y)
	fmt.Println("y after", *y)

}
