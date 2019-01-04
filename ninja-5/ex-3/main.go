package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	ram := truck{
		vehicle: vehicle{
			doors: 2,
			color: "Red",
		},
		fourWheel: true,
	}

	versa := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "Green",
		},
		luxury: false,
	}

	fmt.Println(ram)
	fmt.Println(versa)

	fmt.Println("Ram is", ram.color)
	fmt.Println("Versa is", versa.color)
}
