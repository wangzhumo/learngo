package regex

import (
	"fmt"
	"regexp"
)

const httpUrl = "My email is wangzhumoo@gmail.com@dfjak"

func RunRegexTest() {
	// .*  0个或者多个
	// .+  1个或者多个
	r1 := regexp.MustCompile(".+@.+.\\..+")
	r2 := regexp.MustCompile(`[a-zA-Z0-9]+@.+.\.[a-zA-Z0-9]+`)
	match1 := r1.FindString(httpUrl)
	match2 := r2.FindString(httpUrl)
	fmt.Println(match1)
	fmt.Println(match2)
}
