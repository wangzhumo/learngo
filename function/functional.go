package functions

import "wangzhumo.com/learngo/function/adder"

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func RunFunctional() {
	adder.RunAdder()

}
