package main

import "fmt"

func main() {

	xs1 := []string{"james", "bond", "shaken not stirred"}
	xs2 := []string{"miss", "moneypenny", "helloooooo,james"}
	fmt.Println(xs1)
	fmt.Println(xs2)

	xxs := [][]string{xs1, xs2}
	fmt.Println(xxs)

	for i, xs := range xxs {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v  value: %v", j, val)
		}

	}
}
