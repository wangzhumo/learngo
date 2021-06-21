package channel

import (
	"fmt"
	"time"
)

// 创建Worker 并且需要返回 chan
// monokai
func createWorker(taskId int) chan<- int {
	c := make(chan int)
	go workerTask(taskId, c)
	return c
}

func workerTask(taskId int, c chan int) {
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

func bufferedChannel() {
	c := make(chan int, 3)
	go workerTask(510, c)

	c <- 21
	c <- 22
	c <- 23
	c <- 24

	time.Sleep(time.Microsecond)
}

func RunChannel() {
	channelDemo()
	bufferedChannel()

	time.Sleep(time.Second)
}
