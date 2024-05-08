package main

import (
	"fmt"
)

func arrays(){
	var rank [5] int = [5]int{1, 2, 3, 4, 5}
	teams := [5]string{"Reds", "Blue Jays", "Redsox", "Whitesox", "Rays"}
	city := [...]string{"Cincinnati", "Toronto", "Boston", "Chicago", "Tampa Bay"}
	fmt.Println(teams)
	fmt.Println(rank)
	fmt.Println(city)
	fmt.Println(len(city))
	fmt.Println(teams[2])
	teams[2] = "Bats"
	fmt.Println(teams)

	for place := 0; place < len(teams); place++ {
		fmt.Printf("Your team name is the %s\n",teams[place])
	}
	for _, location := range city {
		fmt.Println("You are in", location)
	}

}

func arrays1(){
	fullName := [3][2]string{{"John", "Smith"},{"Frank", "Lee"}, {"Mark", "Morris"}}
	for _, firstLast := range fullName {
		fmt.Printf("Your first name is %s and your last name is %s\n",firstLast[0],firstLast[1])
	}
}

func main(){
	arrays()
	arrays1()

}