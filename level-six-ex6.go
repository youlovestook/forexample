package main

import "fmt"

func main() {

	func() {
		x := 0
		x++
		fmt.Println(x)
	}()

}
