package main

import (
	"fmt"
	"time"
)

func doLogic(structure string, channel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(structure)
		time.Sleep(3 * time.Second)
	}
	channel <- " is done."
}

func main() {
	channel := make(chan string)
	go doLogic("Processing go here wait !", channel)
	msg := <-channel
	fmt.Println(msg)
}
