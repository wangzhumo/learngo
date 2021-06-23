package channel

import (
	"fmt"
	"time"
)

// 每一个任务
type workItem struct {
	in   chan int
	done chan bool
}

// 创建Worker 并且需要返回 chan
func createDoneWorker(taskId int) workItem {
	c := workItem{
		in:   make(chan int),
		done: make(chan bool),
	}
	go workerDoneTask(taskId, c.in, c.done)
	return c
}

// 实际上的任务执行
func workerDoneTask(taskId int, in chan int, done chan bool) {
	// range 也可以知道channel的关闭
	for value := range in {
		fmt.Printf("goroutine id = %d value = %d \n",
			taskId, value)

		// 通知任务已经完成了
		go func() {
			done <- true
		}()
	}

}

// 开启协程执行任务
func channelDoneDemo() {
	var channels [10]workItem
	for i := 0; i < 10; i++ {
		channels[i] = createDoneWorker(i)
	}

	for i, item := range channels {
		item.in <- 3 * i
	}

	for i, item := range channels {
		item.in <- 8 * i
	}

	// 应该是并行的先直接执行,然后统一的获取任务完成
	for _, item := range channels {
		<-item.done
		<-item.done
	}
	// 此时是有问题的，因为发送者也会等待 done 被接收(因为我们重复了两次任务)
	// 就是说  第一个任务发送出去   发出done信号
	//        第二个任务开始执行发送，
	//		  但是因为第一个任务还在等待done被接收
	// 所以，卡住了
}

func RunChannelDone() {
	channelDoneDemo()

	time.Sleep(time.Second)
}
