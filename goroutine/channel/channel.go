package channel

import (
	"fmt"
	"time"
)

func workerTask(taskId int,c chan int) {
	for  {
		fmt.Printf("goroutine id = %d value = %d \n",taskId, <-c)
	}
}

func channelDemo()  {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go workerTask(i,channels[i])
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 3*i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 8*i
	}
}


func RunChannel()  {
	channelDemo()
	time.Sleep(time.Second)
}
