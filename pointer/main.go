package main

import "fmt"

func main() {
	var x int = 10
	var p *int = &x

	fmt.Printf("Value of x: %d\n", x)          // 10
	fmt.Printf("Address of x: %p\n", &x)       // 0xc0000220b0
	fmt.Printf("Value of p: %p\n", p)          // 0xc0000220b0  => variable p holds the address of var x
	fmt.Printf("Value pointed by p: %d\n", *p) // the value that p vold is 10 as x that lead to when i change *p it change x

	*p = 20

	fmt.Printf("Value of x: %d\n", x)          //20
	fmt.Printf("Value pointed by p: %d\n", *p) //20
}
