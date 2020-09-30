package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("should not print")

	case (2 == 5):
		fmt.Println("should not print")

	case (7 > 3):
		fmt.Println("this will print because 7 is greater than 3")

	case (4 == 490):
		fmt.Println("will never get to this point")
	}
}

// generally its a bad idea to  use fallthrough
