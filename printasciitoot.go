package main

import "fmt"

//printing ascii chars

func main() {

	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#U\n", i, i)
	}
}
