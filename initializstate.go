package main

// you can have multiple statements on a line if you seperate with a ;
// however the linter i am using wont allow it

import "fmt"

func main() {

	if x := 42; x == 42 { // initialization statement is in the statement(limiting scope)

		fmt.Println("001")
	}
	fmt.Println("this is a statement")
	fmt.Println("   here is another statement")
}
