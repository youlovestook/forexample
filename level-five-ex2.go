package main

import "fmt"

//struct is a data structure that aggregates values of different type

type person struct {
	first    string
	last     string
	icecream []string
}

func main() {
	p1 := person{
		first:    "bill",
		last:     "frisell",
		icecream: []string{"vanilla", "strawberry"},
	}

	p2 := person{
		first:    "charlie",
		last:     "watts",
		icecream: []string{"chocolate", "caramel"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.first)

		for _, val := range v.icecream {
			fmt.Println(val)
		}

		fmt.Println("-----------")
	}

	//for _, v := range p1.icecream {
	//	fmt.Println("flavor : ", v)
	//}

	//fmt.Println(p2)

}
