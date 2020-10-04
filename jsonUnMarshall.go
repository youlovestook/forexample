package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {

	s := `[{"First":"james","Last":"bond","Age":45},{"First":"miss","Last":"moneypenny","Age":27}]`
	bs := []byte(s)

	fmt.Printf("%T\n", s)  //string
	fmt.Printf("%T\n", bs) //uint8  byte is an alias for uint8
	fmt.Println("oogie")

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("all the data : ", people)

	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
