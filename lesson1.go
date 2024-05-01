package main

import "fmt"

func name(){
	fname:= "Cameron"
	nname:= "Cam"
	fmt.Println("Hello, my name is", fname,"and I like to go by",nname)
	fmt.Printf("Welcome %v! I will call you %v!\n", fname, nname,)
}


func main() {
	name()
}