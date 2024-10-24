package main

import (
	"fmt"
	"main/client"
	"main/server"
	"time"
)

func main() {
	go server.Main()
	time.Sleep(5 * time.Second)
	currentClient := client.NewClient("http://localhost:8080")
	vers, err := currentClient.GetVersion()
	if err == nil {
		fmt.Println(vers)
	} else {
		fmt.Println(err)
	}

	word := "Hello"
	d, err := currentClient.PostDecode(word)
	if err == nil {
		fmt.Println(d)
	} else {
		fmt.Println(err)
	}

	code, status, err := currentClient.GetHardOp()
	if err == nil {
		if status != 0 {
			fmt.Println(code)
		} else {
			fmt.Println(code, status)
		}
	} else {
		fmt.Println(err)
	}

}


