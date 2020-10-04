package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("go go go")
	fmt.Fprintln(os.Stdout, "yo yo yo wassup")
	io.WriteString(os.Stdout, "yo yo yo wassup... again\n")
}
