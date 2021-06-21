package tryDefer

import "fmt"

func tryDefer() {
	defer fmt.Println("defer print ----2")
	defer fmt.Println("defer print ----3")
	fmt.Println("defer after print 1")

	defer fmt.Println("defer print ----4")
}

func RunDefer()  {
	tryDefer()
}