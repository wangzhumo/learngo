package main

import (
	"fmt"

	fileServer "wangzhumo.com/learngo/errorhandler/server"
	httpClient "wangzhumo.com/learngo/http"
)

func main() {
	fmt.Println(">>>>>>>>>>>>>>start golang application")
	fileServer.RunFileServer()
	httpClient.RunHttpClient()
	fmt.Println(">>>>>>>>>>>>>>end golang application")
}
