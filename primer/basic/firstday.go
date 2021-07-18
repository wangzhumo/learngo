package basic

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// 常量，枚举，函数，匿名函数，可变参数

//常量
func constsFeild() {
	const a, b = 3, 4

}

//枚举
//普通枚举
//自增枚举
func enums() {
	const (
		red    = 0
		green  = 1
		yellow = 2
	)

	// iota 自增值
	const (
		colorRed = iota
		_
		colorGreen
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(red, green, yellow)
	fmt.Println(colorRed, colorGreen)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func mathPow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 函数参数
func applyFunc(op func(int, int) int, a, b int) int {
	var pot = reflect.ValueOf(op).Pointer()
	var opName = runtime.FuncForPC(pot).Name()

	fmt.Printf("call fun %s with params (%d , %d) ", opName, a, b)
	return op(a, b)
}

//可变参数
func sumArgs(agrs ...int) int {
	var sum = 0
	for value := range agrs {
		sum += value
	}
	return sum
}

func RunFirstDay() {
	fmt.Println("firstday")

	enums()
	constsFeild()
	fmt.Println(applyFunc(mathPow, 3, 4))
	fmt.Println(sumArgs(1, 2, 3, 4, 5))
	fmt.Println(applyFunc(func(i1, i2 int) int {
		return int(math.Pow(float64(i1), float64(i2)))
	}, 3, 4))

}
