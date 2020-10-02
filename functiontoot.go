package main

import "fmt"

func main() {

	fmt.Println("")
	foo()
	bar(" you big dummy")
	s1 := woo("bill frisell")
	fmt.Println(s1)
	fmt.Println("---------------------")
	x, y := cootie("bob", "ross")
	fmt.Println(x)
	fmt.Println(y)

}

// func (r reciever) identifier (parameters) (return(s)) {...}
// need to know the difference between paramaters and arguments
// we call with params and we pass in arguments

func foo() {
	fmt.Println("this is foo")
}

// everything in go is PASS BY VALUE

func bar(s string) {
	fmt.Println("this is bar", s)

}

func woo(s string) string {
	return fmt.Sprint("Hello from woo  ", s)
}

func cootie(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, "  says hello")
	b := false
	return a, b

}
