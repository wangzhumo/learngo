package recovers

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			// 是一个error,并获取到了panic的数据
			fmt.Printf("Error find : %s \n", err)
		} else {
			panic(r)
		}
	}()
	// panic Error
	// panic(errors.New("Error start "))

	// 无法处理，recover会再次转发panic
	panic(123)

}

func RunRecovers() {
	tryRecover()
}
