package main

import "fmt"

func birthday(pointerAge *int) {
	if *pointerAge > 150 {
		panic("The age is old to be true !")
	}
	fmt.Printf("The Pointer is %v and the value is %v \n", pointerAge, *pointerAge)
	*pointerAge++
}

func main() {
	age := 160
	birthday(&age)
	fmt.Println(age)

}
