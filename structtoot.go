package main

import "fmt"

//struct is a data structure that aggregates values of different type

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "bill",
		last:  "frisell",
		age:   27,
	}

	p2 := person{
		first: "oscar",
		last:  "delayhoya",
	}

	fmt.Println(p1.first, " last: ", p1.age)
	fmt.Println(p2)

	fmt.Println("ya")
}
