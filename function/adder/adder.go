package adder

import "fmt"


// 实现斐波那也数列
// 1 1 2 3 5 8 13
func fibonacci() func() int {
	// 定义前两个数的保存地方
	a,b := 0,1
	return func() int {
		a,b = b,a+b
		return a
	}
}



func adder() func(int) int {
	sum := 0
	fmt.Println(sum)
	return func(v int) int {
		fmt.Printf("%d += %d \n", sum, v)
		sum += v
		return sum
	}
}

func RunAdder() {
	var a = adder()
	for i := 0; i < 10; i++ {
		a(i)
		//fmt.Printf("++ %d = %d \n", i, a(i))
	}

	f := fibonacci()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
