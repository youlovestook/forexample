package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {

	fmt.Println("ya")

	p1 := person{
		first: "bill",
		last:  "frisell",
		age:   57,
	}

	speak(p1)

}

func speak(p person) {
	fmt.Println("Can you hear me now i am :", p.first, " ", p.last)
}
