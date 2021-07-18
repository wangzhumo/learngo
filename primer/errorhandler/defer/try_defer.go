package tryDefer

import (
	"bufio"
	"fmt"
	"os"
	functions "wangzhumo.com/learngo/function"
)

// panic
// 1.停止当前的函数
// 2.向上返回，执行每一层的defer语句
// 3.如果没有recover,程序退出

// recover
// 1.仅仅在defer调用中是使用
// 2.获取panic

// 需要注意的是  defer  先进后出
// defer after print 1
// defer print ----4
// defer print ----3
// defer print ----2
func tryDefer() {
	defer fmt.Println("defer print ----2")
	defer fmt.Println("defer print ----3")
	fmt.Println("defer after print 1")

	defer fmt.Println("defer print ----4")
}

func writeFile(fileName string) {
	// create file with name
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	// close file
	defer file.Close()

	// write with buffer.io
	writer := bufio.NewWriter(file)
	// flush context
	defer writer.Flush()

	f := functions.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func writeFileError(fileName string) {
	// open file with name
	file, err := os.OpenFile(fileName, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		// 处理这个异常
		if pathError, ok := err.(*os.PathError); !ok {
			// 已知返回err是 *PathError ，转换失败就无法继续处理
			panic(err)
		} else {
			fmt.Println(pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}

	// close file
	defer file.Close()

	// write with buffer.io
	writer := bufio.NewWriter(file)
	// flush context
	defer writer.Flush()

	f := functions.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func RunDefer() {
	tryDefer()
	//writeFile("errorhandler/defer/fib.txt")
	writeFileError("errorhandler/defer/fib.txt")
}
