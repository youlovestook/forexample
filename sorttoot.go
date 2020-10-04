package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr.No"}

	fmt.Println(xi)
	fmt.Println(xs)

	sort.Ints(xi)
	fmt.Println("sorted :", xi)
	fmt.Println("-------------------------------")
	sort.Strings(xs)
	fmt.Println(xs)

}
