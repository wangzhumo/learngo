package basic

import "fmt"

// 他的作用域在这个包下面   package main
// var ()  集中定义一堆变量
var (
	ga = 1
	gb = "dddd"
	gc = false
)

func RunBasic() {
	fmt.Println("Hello World")
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()

	fmt.Println(ga, gb, gc)
}

// 打印零值
// var {name} type
func variableZeroValue() {
	var a int
	var b string
	fmt.Printf("%d ,%q\n", a, b)
}

// 初始值
// var {name1} {name2} type = value1,value2
func variableInitValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

// 类型推断
// var {name1} {name2} = value1,value2
// {name1} {name2} := value1,value2  (:=  只能在函数内使用)
func variableTypeDeduction() {
	// 依次复制
	var a, b, c = 3, 5, "def"
	d, e, f := 5, "4", true
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
}

// 常量的定义
// const 的数值可以作为各种类型使用
func constValues() {
	const fileName = "down.txt"
	const (
		a = 1
		b = "text"
	)
}

// 枚举的使用
func enumUse() {
	const (
		cpp    = 1
		java   = 2
		dart   = 3
		golang = 4
	)

	// 1 2 3 4

	// 这里使用的 iota 表示自增
	// _ 表示这里跳过  0 2 3 4
	const (
		js = iota
		_
		python
		rust
	)

	// 0 2 3 4

}

// byte   rune(这个就是char  32位的 4字节)
// float32 float64 complex64 complex128
// int int8 int16 int32 int 64 uintptr (还有无符号的）
// bool
// string
