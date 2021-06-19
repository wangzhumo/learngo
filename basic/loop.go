package basic

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunLoop() {
	first := 13
	fmt.Println(convertToBinary(first))
	readFileByLine("README.md")
}

// 把数字转为2进制
// 其中num 不需要初始化，则直接 _ ; 条件 ; 操作
func convertToBinary(num int) string {
	var result = ""
	for ; num > 0; num /= 2 {
		// 如果除以2
		temp := num % 2
		result = strconv.Itoa(temp) + result
	}
	return result
}

// 按照行读取file中的数据
// 在这个for中，完全可以用while替代，这里只提供了循环的条件
// [死循环] 如果连这个条件都不提供，就成了死循环
func readFileByLine(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		//报错
		panic(err)
	}
	// 行读取
	scanner := bufio.NewScanner(file)
	// 拿到每一行
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
