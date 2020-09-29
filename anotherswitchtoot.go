package main

import "fmt"

func main() {

	switch "Bond" {
	case "MoneyPenny":
		fmt.Println("miss money")
	case "Bond":
		fmt.Println("bond james")
	case "Q":
		fmt.Println("this is Q")
	default:
		fmt.Println("this is default")
	}
}
