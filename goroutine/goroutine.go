package goroutines

import (
	"fmt"
	"time"
)

// 打印数据
func goroutine() {
	for i := 0; i < 10; i++ {
		go func(value int) {
			// 此时打印
			fmt.Println(value)
		}(i)
	}
}

// 使用channel
func goChannelDemo() {
	channels := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func(value int, c chan bool) {
			fmt.Println(value)
			c <- true
		}(i, channels)
	}

	for j := 0; j < 10; j++ {
		<-channels
		//fmt.Println("channel",j)
	}

	close(channels)
	fmt.Println("channel done")
}

// 使用select
func goSelectDemo() {
	// 创建两个channel
	c1 := make(chan int, 20)
	c2 := make(chan string, 20)
	// 执行
	go func(c chan int) {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}(c1)

	c2 <- "a"
	c2 <- "b"

	go func() {
		for {
			select {
			case v1 := <-c1:
				fmt.Println("get c1 = ", v1)
			case v2 := <-c2:
				fmt.Println("get c2 = ", v2)
			default:
				fmt.Println("...")
			}
		}
	}()
}

// 模拟任务池.

func RunGoroutine() {
	goSelectDemo()
	time.Sleep(1 * time.Second)
}
