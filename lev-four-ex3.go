package main

import "fmt"

func main() {

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	//fmt.Println(x)

	//for i, v := range x {
	//	fmt.Println(i, v)
	//}

	y := x[0:5]
	z := x[5:10]
	aa := x[2:7]
	bb := x[1:6]
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(aa)
	fmt.Println(bb)
	//fmt.Printf("%T\n", x)
}
