package main

import "fmt"

// maps -> key value store
// key is the indicator value is the value
// its an unordered list

func main() {

	m := map[string]int{
		"james":      32,
		"moneypenny": 27,
	}
	fmt.Println(m)

	fmt.Println(m["james"])

	fmt.Println(m["chuckie"]) // returns 0 because not in the map

	v, ok := m["chuckie"] // comma okay idiom
	fmt.Println(v)
	fmt.Println(ok)

	// below we combine with if statement... very idiomatic chunk of code in go
	// basically if this is in the map then do this
	if v, ok := m["james"]; ok {
		fmt.Println("this is the if print", v)
	}

}
