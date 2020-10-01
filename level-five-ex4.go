package main

// anonymous struct is a struct without a name
// if you only need a struct to use in a limited area
// you can use an anonymous struct

import "fmt"

func main() {
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "james",
		friends: map[string]int{
			"hank": 555,
			"carl": 776,
			"phil": 888,
		},
		favDrinks: []string{
			"martini",
			"beer",
			"blood",
		},
	}

	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for _, v := range s.favDrinks {
		fmt.Println(v)
	}

}
