package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	favIceCream []string
}

func main() {
	persons := make(map[string]person)

	persons["Cool"] = person{
		"Joe",
		"Cool",
		[]string{"Chocolate", "Rocky Road", "Double Fudge Brownie"},
	}

	persons["Van Pelt"] = person{
		"Lucy",
		"Van Pelt",
		[]string{"Vanilla", "Strawberry", "Peanut Butter"},
	}

	for _, v := range persons {
		fmt.Printf("%s %s's favorite ice cream flavors:\n", v.firstName, v.lastName)
		for _, flavor := range v.favIceCream {
			fmt.Printf("\t%s\n", flavor)
		}
	}
}
