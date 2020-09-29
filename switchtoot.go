package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("should not print")

	case (2 == 2):
		fmt.Println("should not print")
		fallthrough
	case (7 == 3):
		fmt.Println("not true!!!")
		fallthrough // allows following statements to be run even if true
	case (4 == 4):
		fmt.Println("should not print because when it finds a true case it exits (unless fallthrough)")
	}

}

// generally its a bad idea to  use fallthrough
