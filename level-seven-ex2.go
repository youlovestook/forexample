package main

import "fmt"

type person struct {
	name string
}

func main() {

	fmt.Println("ya ya ya")

	p1 := person{
		name: "james bond",
	}

	fmt.Println(p1.name)
	changeMe(&p1)
	fmt.Println(p1.name)
}

func changeMe(p *person) {
	//p.name = "miss moneypenny"
	(*p).name = "Miss Moneyp"
}
