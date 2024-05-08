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

func slice(){
	fruits := []string{"Apple", "Orange", "Kiwi", "Blueberry"}
	fmt.Println(fruits)
	var num [5] int = [5]int{1, 2, 3, 4, 5}
	sliced := num[1:4]
	subSlice := sliced[0:2]
	fmt.Println(sliced)
	fmt.Println(subSlice)
	numSlice := make([]int, 5, 10)
	fmt.Println(numSlice)
	sliced = append(sliced, 10, 12, 14)
	fmt.Println(sliced)
	sliced = append(sliced, subSlice...)
	fmt.Println(sliced)

}

func main(){
	slice()

}