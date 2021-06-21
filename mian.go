package main

import (
	"fmt"
	"wangzhumo.com/learngo/goroutine/channel"
)

func main() {
	fmt.Println(">>>>>>>>>>>>>>start golang application")
	channel.RunChannel()
	fmt.Println(">>>>>>>>>>>>>>end golang application")
}
