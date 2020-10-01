package main

import "fmt"

//this is an example of an embedded struct

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "bill",
			last:  "frisell",
			age:   50,
		},
		ltk: true,
	}

	p2 := person{
		first: "oscar",
		last:  "delayhoya",
	}

	fmt.Println(sa1.first, " last: ", sa1.age, sa1.ltk)
	fmt.Println(p2)

	fmt.Println("ya")
}
