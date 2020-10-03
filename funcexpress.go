package main

import "fmt"

func main() {

	fmt.Println("lovey")

	f := func() {
		fmt.Println("my first func expression")
	}
	f()

	f2 := func(x int) {
		fmt.Println("Jackie Robinson wore number: ")
	}
	f2(42)
}
