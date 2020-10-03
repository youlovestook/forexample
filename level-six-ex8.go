package main

import "fmt"

func main() {

	fmt.Println("ya ya ya")
	x := getOneBack()
	fmt.Println(x)
	i := x()
	fmt.Println(i)
}

func getOneBack() func() int {
	return func() int {
		return 451
	}
}
