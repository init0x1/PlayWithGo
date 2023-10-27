package main

import "fmt"

func main() {
	var operation string
	var num1, num2 int
	fmt.Println("Enter the first number: ")
	fmt.Scan(&num1)
	fmt.Println("Enter the second number: ")
	fmt.Scan(&num2)
	fmt.Println("Choose an operation (+,-,*,/):")
	fmt.Scan(&operation)

	switch operation {
	case "+":
		fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%d - %d = %d\n", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%d * %d = %d\n", num1, num2, num1*num2)
	case "/":
		if num2 == 0 {
			panic("you cant divide over zero")
		}
		fmt.Printf("%d / %d = %d\n", num1, num2, num1/num2)
	}
}
