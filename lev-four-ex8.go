package main

import "fmt"

func main() {

	m := map[string][]string{
		"bond_james":      []string{"shaken, not stirred", "Martinis", "women"},
		"miss_moneypenny": []string{"ice", "bourbon", "cheese"},
		"bill_frisell":    []string{"guitars", "coffee", "history"},
	}
	//fmt.Println(m)

	for k, v := range m {
		fmt.Println("this is the record for, ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}

	}

}
