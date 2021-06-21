package basic

import "fmt"

// int 是值传递
func printArray2(arr *[5]int) {
	arr[0] = 90
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printSliceArray(arr []int) {
	arr[0] = 989
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func RunArrays() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{2, 4, 6, 8, 12}

	//二维数组
	var grid [4][4]int

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(grid)

	// 遍历数组
	for i := range arr3 {
		fmt.Println("value = ", i)
	}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	for _, v := range arr3 {
		fmt.Println(v)
	}

	printArray2(&arr3)
	fmt.Println(arr3)
	// 切片
	fmt.Println("Slice 1")
	s1 := arr3[1:3]
	fmt.Println(s1)
	fmt.Printf("s1 = %v , len = %d , cap = %d \n", s1, len(s1), cap(s1))

	// 切片可以向后扩展
	// 但是不可以向前扩展
	fmt.Println("Slice 2")
	s2 := s1[2:4]
	fmt.Println(s2)

}
