package channel

func channelDemo()  {
	c := make(chan int)
	c<-1
}


func RunChannel()  {
	channelDemo()
}
