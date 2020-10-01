package main

import "fmt"

//struct is a data structure that aggregates values of different type

type person struct {
	first    string
	last     string
	icecream []string
}

func main() {
	p1 := person{
		first:    "bill",
		last:     "frisell",
		icecream: []string{"vanilla", "strawberry"},
	}

	p2 := person{
		first:    "charlie",
		last:     "watts",
		icecream: []string{"chocolate", "caramel"},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)

	for _, v := range p1.icecream {
		fmt.Println("flavor : ", v)
	}

	fmt.Println(p2)

}
