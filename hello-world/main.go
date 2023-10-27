package main

import (
	"fmt"
	"hello/world/users"
)

func birthday(pointerAge *int) {
	if *pointerAge > 150 {
		panic("The age is old to be true !")
	}
	fmt.Printf("The Pointer is %v and the value is %v \n", pointerAge, *pointerAge)
	*pointerAge++
}

func main() {
	age := 100
	birthday(&age)
	fmt.Println(age)

	user1 := users.User{
		UserID:   "100",
		UserName: "init0x1",
		Password: "letmein",
	}

	message := user1.CreateUser()
	fmt.Println(message)
}
