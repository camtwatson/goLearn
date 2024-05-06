// operators 
package main 
import "fmt"

func comparison() {
	county:= "Jefferson"
	county2:= "Bullet"
	var num1 int
	var num2 int

	fmt.Println("Are the county's the same?", county == county2)
	fmt.Println("Enter a number")
	fmt.Scanf("%v", &num1)
	fmt.Println("Enter another number")
	fmt.Scanf("%v", &num2)
	fmt.Println("The first number is larger than the second number:",num1 > num2)

}

func logical() {
	var num3 int
	fmt.Println("Please enter a number:")
	fmt.Scanf("%v", &num3)
	fmt.Println("Is your number greater than 5 and less than 500:", (num3 > 5) && (num3 <500))
	fmt.Println("Your number is less than 5 or greater than 500:", (num3 < 5) || (num3 > 500))
	fmt.Println("Your number is not equal to 0:", (num3 != 0))

}

func loop1() {
	var yourAge int
	fmt.Println("How old are you?")
	fmt.Scanf("%v", &yourAge)
	yearWait:= (16 - yourAge)

	if yourAge >= 18 {
		fmt.Println("You are old enough to drive")
	} else if yourAge == 16 || yourAge == 17 {
		fmt.Println("You can get your drivers permit")
	} else {
		fmt.Printf("You have to wait %v years to drive\n", yearWait)
	}

}


func main() {
	loop1()
}
