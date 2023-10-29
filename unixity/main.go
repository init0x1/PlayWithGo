package main

import (
	"fmt"

	"github.com/0xAbdoAli/PlayWithGo/unixity/models"
)

func main() {
	// we can use Service like this way

	backend := models.Service{ServiceId: "fds4da845sd", ServiceName: "Backend", TeamMembers: []string{"init0x1", "0xCrypt00o"}}

	//Or

	pentrationTesting := models.NewService("h57sdre85s", "Pentesting", []string{"init0x1", "0xCrypt00o"})

	fmt.Println(backend.ListServices())
	fmt.Println(pentrationTesting.ListServices())

	//use ServiceInfo interface
	fmt.Println("================   use ServiceInfo interface  ===============")
	pentrationTesting.DisplayServiceInfo()

}
