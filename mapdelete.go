package main

import "fmt"

// delete(<map name>, "key")
// if you delete a key that doesnt exist no error is thrown

func main() {
	m := map[string]int{
		"james":      32,
		"moneypenny": 27,
	}
	fmt.Println(m)

	delete(m, "moneypenny")

	fmt.Println(m)

	if v, ok := m["james"]; ok {
		fmt.Println("james is in there", v)
	}
}
