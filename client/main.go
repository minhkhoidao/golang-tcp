package main

import (
	"fmt"
	"golang-tcp/tcp"
)

func main() {
	text := "Hello, this is a test."
	err := tcp.SendTextToServer(text)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
