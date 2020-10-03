package main

import "fmt"

// a value can be of more than one type

type hotdog int

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

type human interface {
	speak()
}

// interfaces allow us to define behaviour
// fun (r reciever) identifier(parameters) (return(s)) {code here}
// a value can be of more than one type

func bar(h human) {
	fmt.Println("i was passed into bar,", h)
}

func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last)
}

func (p person) speak() {
	fmt.Println("the person speak")
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Bill",
			"Frisell",
		},
		ltk: true,
	}

	p1 := person{
		first: "dr.",
		last:  "yes",
	}

	x := person{
		first: "carl",
		last:  "gerbschmidt",
	}

	bar(sa1)
	bar(sa2)
	bar(p1)
	bar(x)

	fmt.Println("---------------------------------")
	// conversion
	var b hotdog = 42
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	var y int
	y = int(b)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	//assertion

}
