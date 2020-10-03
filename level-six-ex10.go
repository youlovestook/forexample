package main

import "fmt"

// closure is limiting the scope of a variable
// limiting scope is very important

var cola int // package level scope

func main() {
	foo()
	fmt.Println(cola)
	{

		soda := "coca"
		fmt.Println(soda)
	}

	//fmt.Println(yyz)  this will not print because the scope of yyz is only in the closure

}

func foo() {
	fmt.Println("hello")
	cola++
}
