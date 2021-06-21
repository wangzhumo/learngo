package basic

import (
	"fmt"
	"io/ioutil"
)

func RunBranch() {
	// readFile("")
	// The system cannot find the file specified.
	readFileLikeGo("README.md")
	showYouSwitch(100)
}

func readFile(fileName string) {
	context, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", context)
	}
}

// if 的条件中可以赋值
// ;  可以有多个条件
func readFileLikeGo(fileName string) {
	if context, err := ioutil.ReadFile(fileName); err != nil {
		fmt.Println("error ", err)
	} else {
		fmt.Println(string(context))
	}
}

// switch
// switch 自动break ， 如果需要向下就使用fallthrough
func showYouSwitch(score int) {

	// 只能直接比较
	switch score {
	case 60:
		fallthrough
	case 66:
		fmt.Println(score)
	case 100:
		fmt.Println("100")
	default:

	}

	// 可以加入条件语句
	switch {
	case score < 0 || score > 100:
		// 报错
		panic(fmt.Sprintf("wrong score: %d", score))
	case score < 60:
		fmt.Println("不及格")
	case score < 90:
		fmt.Println("优良")
	case score <= 100:
		fmt.Println("优秀")
	}

}
