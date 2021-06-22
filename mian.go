package main

import (
	"fmt"

	"wangzhumo.com/learngo/crawler"
)

func main() {
	fmt.Println(">>>>>>>>>>>>>>start golang application")
	crawler.RunCrawlerTask()
	fmt.Println(">>>>>>>>>>>>>>end golang application")
}
