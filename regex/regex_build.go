package regex

import (
	"fmt"
	"regexp"
)

const httpUrl = "My email is wangzhumoo@gmail.com"

func RunRegexTest() {
	r := regexp.MustCompile("wangzhumoo@gmail.com")
	match := r.FindString(httpUrl)
	fmt.Println(match)
}
