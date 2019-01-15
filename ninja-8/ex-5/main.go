package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type users []user

func (ul users) Len() int {
	return len(ul)
}

func (ul users) Less(i, j int) bool {
	return ((ul[i].Age < ul[j].Age) || ((ul[i].Age == ul[j].Age) && ul[i].Last < ul[j].Last))
}

func (ul users) Swap(i, j int) {
	ul[i], ul[j] = ul[j], ul[i]
}

func printUser(u user) {
	fmt.Printf("%s %s (%d)\n", u.First, u.Last, u.Age)
	fmt.Println("Sayings:")
	for _, saying := range u.Sayings {
		fmt.Printf("\t%s\n", saying)
	}
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	u4 := user{
		First: "Auric",
		Last:  "Goldfinger",
		Age:   32,
		Sayings: []string{
			"No, Mr. Bond. I expect you to die!",
			"You are quite right, Mr Bond. You are worth more to me alive.",
			"Two holes to go.",
		},
	}

	userSet := users{u1, u2, u3, u4}

	sort.Sort(userSet)

	for _, u := range userSet {
		sort.Strings(u.Sayings)
	}

	for _, u := range userSet {
		fmt.Println("------------------")
		printUser(u)
	}

}
