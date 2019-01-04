package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	favIceCream []string
}

func main() {
	joe := person{
		"Joe",
		"Cool",
		[]string{"Chocolate", "Rocky Road", "Double Fudge Brownie"},
	}

	lucy := person{
		"Lucy",
		"Van Pelt",
		[]string{"Vanilla", "Strawberry", "Peanut Butter"},
	}

	fmt.Println(joe)
	fmt.Println(lucy)

	fmt.Printf("%s %s's favorite ice cream flavors:\n", joe.firstName, joe.lastName)
	for _, flavor := range joe.favIceCream {
		fmt.Printf("\t%s\n", flavor)
	}

	fmt.Printf("%s %s's favorite ice cream flavors:\n", lucy.firstName, lucy.lastName)
	for _, flavor := range lucy.favIceCream {
		fmt.Printf("\t%s\n", flavor)
	}
}
