package interfaces

import (
	"fmt"
	"time"
	"wangzhumo.com/learngo/interface/impl"
)

const requestHttpUrl = "https://www.baidu.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

//RetrieverAndPoster 接口的组合
type RetrieverAndPoster interface {
	Retriever
	Poster
}

func download(retriever Retriever) string {
	return retriever.Get(requestHttpUrl)
}

func post(p Poster) string {
	return p.Post(requestHttpUrl, map[string]string{
		"UserAgent": "OkHttp3.9.10/Wangzhumo-PC",
		"name":      "wangzhumo",
	})
}

func combine(s RetrieverAndPoster) string {
	s.Post(requestHttpUrl, map[string]string{
		"UserAgent": "OkHttp3.9.11/Wangzhumo-PC",
		"name":      "wangzhumo-pc",
	})
	return s.Get(requestHttpUrl)
}

func showType(r Retriever) {
	if t, ok := r.(*impls.Download); ok {
		fmt.Println("UserAgent:", t.UserAgent)
		fmt.Println("TimeOut:", t.TimeOut)
	}
}
func RunInterface() {
	d := &impls.Download{
		UserAgent: "OkHttp3.9.0/Android",
		TimeOut:   time.Minute,
	}

	showType(d)

	q := impls.Queue{1}
	q.Push(123)
	q.Push(321)
	q.Push(99)

	fmt.Println("isEmpty = ", q.IsEmpty())
	fmt.Println(q)
	down := &impls.Download{}
	combine(down)
	showType(down)
	fmt.Println(q.String())
}
