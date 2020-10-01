package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourwheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourwheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: true,
	}
	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.doors)
}
