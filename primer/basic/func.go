package basic

import (
	"fmt"
	"reflect"
	"runtime"
)

// return 两个值
func div(a, b int) (q, r int) {
	return a / b, a * b
}

func div2(a, b int) (q, r int) {
	q = a / b
	r = a * b
	return q, r
}

func add(a, b int) int {
	return a + b
}

// 函数做参数
func apply(op func(int, int) int, a, b int) int {
	// 获取到这个方法的指针
	pOp := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(pOp).Name()
	fmt.Printf("call function %s with (%d,%d) \n", opName, a, b)
	return op(a, b)
}

// 可变参数,没有默认参数
func sum(numbers ...int) int {
	sum := 0
	for index := range numbers {
		sum += numbers[index]
	}
	return sum
}

func RunFunc() {
	q, r := div(1, 3)
	fmt.Println(q, r)
	p, _ := div2(2, 3)
	fmt.Println(p)
	result := apply(add, 2, 3)

	fmt.Println(result)
}
