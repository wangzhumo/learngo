package basic

import (
	"fmt"
)

// map
func maps() {
	var m = map[string]string{"name": "wzngzhumo", "age": "18"}
	var m2 = make(map[string]string)
	fmt.Println(m)
	fmt.Println(m2)

	// 取一个值
	fmt.Println(m["name"])

	// 判断是否有key
	var name, err = m["nane"]
	fmt.Println(name, err)

	// 移除一个值
	delete(m, "name")
	fmt.Println(m)
}

// map 遍历
func mapsForeach() {
	var m = map[string]string{"name": "wzngzhumo", "age": "18"}
	for k, v := range m {
		fmt.Printf("key = %s , value = %s   ", k, v)
	}
}

// 数组
func arrays() {
	var arr1 [5]int
	var arr2 = [3]int{3, 4, 5}
	var arr3 = [...]int{2, 4, 5, 6}
	// 二维
	var arr4 [4][4]int
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	printArray(arr1)
	fmt.Println(arr1)
	printArraySlice(arr1[:])
	fmt.Println(arr1)
}

// 数组遍历
func arraysForeach() {
	var arr = [...]int{2, 4, 5, 6, 7, 22, 321, 32, 8}
	for i, v := range arr {
		fmt.Print(i, v, "  ")
	}
	fmt.Println("")

	for _, v := range arr {
		fmt.Print(v)
	}
}

// 这一种数组的打印，但是是值的引用
func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Print(i, v, "  ")
	}
	arr[0] = 100
	fmt.Println("")
}

// 这一种是数组的打印，这样的话是切片，
// 这样的话，就是地址的引用了
func printArraySlice(arr []int) {
	for i, v := range arr {
		fmt.Print(i, v, "  ")
	}
	arr[0] = 100
	fmt.Println("")
}

func RunCollection() {
	fmt.Println("firstday_collections")
	//arrays()
	//arraysForeach()

	maps()
	mapsForeach()
}
