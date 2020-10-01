package main

// anonymous struct is a struct without a name
// if you only need a struct to use in a limited area
// you can use an anonymous struct

import "fmt"

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "james",
		last:  "bond",
		age:   32,
	}

	fmt.Println(p1)

}
