package main

import "fmt"

func main() {
	bd := 10
	for bd <= 100 {
		if bd%4 == 0 {
			fmt.Println(bd)
		}

		bd++
	}
}
