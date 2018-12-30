package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	m["q"] = []string{"Cars", "Gadgets", "Electronics"}

	for k, v := range m {
		fmt.Println("Key: ", k)
		for i, t := range v {
			fmt.Printf("\t %d\t%s\n", i, t)
		}
	}
}