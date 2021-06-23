package channel

import (
	"fmt"
	"sync"
	"time"
)

// 每一个任务
type worker struct {
	in chan int
	// 使用外部的WaitGroup
	done func()
}

// 创建Worker 并且需要返回 chan
func createWaitWorker(taskId int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go workerWaitTask(taskId, w)
	return w
}

// 实际上的任务执行
func workerWaitTask(taskId int, w worker) {
	// range 也可以知道channel的关闭
	for value := range w.in {
		fmt.Printf("goroutine id = %d value = %d \n",
			taskId, value)

		// 通知任务已经完成了
		w.done()
	}

}

// 开启协程执行任务
func channelWaitDemo() {
	// 通过waitGroup来做等待任务完成
	var wg sync.WaitGroup
	// 我们会发20次任务
	wg.Add(20)

	var channels [10]worker
	for i := 0; i < 10; i++ {
		channels[i] = createWaitWorker(i, &wg)
	}

	for i, item := range channels {
		item.in <- 3 * i
	}
	for i, item := range channels {
		item.in <- 8 * i
	}
	wg.Wait()
}

func RunChannelWait() {
	channelWaitDemo()

	time.Sleep(time.Second)
}
