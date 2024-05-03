// operators 
package main 
import "fmt"

func comparison() {
	county:= "Jefferson"
	county2:= "Bullet"
	var num1 int
	var num2 int

	fmt.Println(county == county2)
	fmt.Println("Enter a number")
	fmt.Scanf("%v", &num1)
	fmt.Println("Enter another number")
	fmt.Scanf("%v", &num2)
	fmt.Println("The first number is larger than the second number:\n",num1 > num2)


}

func main() {
	comparison()
}
