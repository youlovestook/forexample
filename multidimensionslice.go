package main

import "fmt"

func main() {

	jb := []string{"james", "bond", "choco", "martini"}
	fmt.Println(jb)
	mp := []string{"miss", "moneypenny", "strawberry", "hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}

	fmt.Println(xp)

}
