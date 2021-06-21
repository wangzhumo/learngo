package goroutines

import (
	"fmt"
	"time"
)

func goroutine(n int) {
	var a [10]int
	for i := 0; i < n; i++ {
		go func() {
			for {
				a[0]++
			}
		}()
	}
	fmt.Println(a)
}

func RunGoroutine() {
	goroutine(8)
	time.Sleep(time.Second)
}
