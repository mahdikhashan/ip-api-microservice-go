package main

import (
	"fmt"
	"service/ip"
)

func main() {
	message := ip.IpHello("mahdi")
	fmt.Println(message)
}
