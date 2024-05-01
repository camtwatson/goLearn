package main

import "fmt"

func name(){
	fname:= "Cameron"
	nname:= "Cam"
	fmt.Println("Hello, my name is", fname,"and I like to go by",nname)
	fmt.Printf("Welcome %v! I will call you %v!\n", fname, nname,)
}

func input(){
	city:= ""
	var age int
	fmt.Print("What city were you born in?")
	fmt.Scanf("%s", &city)
	fmt.Print("How old are you?")
	fmt.Scanf("%v", &age)
	fmt.Printf("You are %v and you were born in %s!\n", age, city)

}

func constants(){
	const Country = "United States"
	fmt.Printf("You are in the %v\n", Country)
}

func main() {
	name()
	input()
	constants()
}