package main

import "fmt"

// delete(<map name>, "key")

func main() {
	m := map[string]int{
		"james":      32,
		"moneypenny": 27,
	}
	fmt.Println(m)

	delete(m, "moneypenny")

	fmt.Println(m)
}
