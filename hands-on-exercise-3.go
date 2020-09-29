package main

import "fmt"

var xo int = 42
var yo string = "Bill Frissell"
var zo bool = true

func main() {

	s := fmt.Sprintf("%v\t%v\t%v", xo, yo, zo)

	fmt.Println(s)

}
