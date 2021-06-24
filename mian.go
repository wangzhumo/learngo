package main

import (
	"fmt"

	"wangzhumo.com/learngo/crawler"
	"wangzhumo.com/learngo/regex"
)

func main() {
	fmt.Println(">>>>>>>>>>>>>>start golang application")
	crawler.RunCrawlerTask()
	regex.RunRegexTest()
	fmt.Println(">>>>>>>>>>>>>>end golang application")
}
