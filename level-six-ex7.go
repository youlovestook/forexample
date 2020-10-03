package main

import "fmt"

func main() {
	fmt.Println("honey")

	f2 := func(x int) {
		fmt.Println("Jackie Robinson wore number: ", x)
	}
	f2(42)
}
