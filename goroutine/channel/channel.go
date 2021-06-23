package channel

import (
	"fmt"
	"time"
)

// 创建Worker 并且需要返回 chan
func createWorker(taskId int) chan<- int {
	c := make(chan int)
	go workerTask(taskId, c)
	return c
}

// 实际上的任务
func workerTask(taskId int, c chan int) {
	// range 也可以知道channel的关闭
	for value := range c {
		fmt.Printf("goroutine id = %d value = %d \n",
			taskId, value)
	}
}

func channelDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 3 * i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 8 * i
	}

	time.Sleep(time.Microsecond)
}

// channel的缓冲区
func bufferedChannel() {
	c := make(chan int, 3)
	go workerTask(510, c)

	c <- 21
	c <- 22
	c <- 23
	c <- 24
	c <- 25

	time.Sleep(time.Microsecond)
}

// 当channel关闭后，如果还在收数据，那就是channel的 零值
func channelClose() {

	c := make(chan int)
	go workerTask(51, c)

	c <- 211
	c <- 221
	c <- 231
	c <- 241
	c <- 251

	close(c)

	time.Sleep(time.Microsecond)
}

func RunChannel() {
	channelDemo()
	bufferedChannel()
	channelClose()
	time.Sleep(time.Second)
}
