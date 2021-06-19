package interfaces

import (
	"fmt"
	"time"

	impls "wangzhumo.com/learngo/interface/impl"
)

type Retriever interface {
	Get(url string) string
}

func download(retiever Retriever) string {
	return retiever.Get("https://www.baidu.com")
}

func RunInterface() {
	d := impls.Download{
		UserAgent: "OkHttp3.9.0/Android",
		TimeOut:   time.Minute,
	}
	fmt.Printf("%T", d)
}
