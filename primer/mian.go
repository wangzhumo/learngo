package main

import (
	"fmt"
	goroutines "wangzhumo.com/learngo/goroutine"
)

func main() {
	fmt.Println(">>>>>>>>>>>>>>start golang application")
	goroutines.RunGoroutine()
	fmt.Println(">>>>>>>>>>>>>>end golang application")
}
