package main

import "fmt"

// using an if to do a for loop
// is the format below

func main() {

	x := 1
	for {
		if x > 9 {
			fmt.Println("breaking")
			break

		}
		fmt.Println(x)
		x++
	}

	fmt.Println("done with the loop")
}
