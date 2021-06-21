package main

import (
	"fmt"
	fileServer "wangzhumo.com/learngo/errorhandler/server"
)

func main() {
	fmt.Println(">>>>>>>>>>>>>>start golang application")
	fileServer.RunFileServer()
	fmt.Println(">>>>>>>>>>>>>>end golang application")
}
