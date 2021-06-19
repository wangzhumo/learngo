package interfaces

import (
	"fmt"
	"time"
	"wangzhumo.com/learngo/interface/impl"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,form map[string]string) string
}

type RetrieverAndPoster interface {
	Retriever
	Poster

}

func download(retriever Retriever) string {
	return retriever.Get("https://www.baidu.com")
}

func post(p Poster) string {
	return p.Post("https://www.baidu.com", map[string]string{
		"name":"wangzhumo",
		"age":"18",
	})
}

func combint(s RetrieverAndPoster)  {
	s.Get()
	s.Post()
}


func showType(r Retriever)  {
	if t, ok := r.(*impls.Download); ok {
		fmt.Println("UserAgent:",t.UserAgent)
		fmt.Println("TimeOut:",t.TimeOut)
	}
}
func RunInterface() {
	d := &impls.Download{
		UserAgent: "OkHttp3.9.0/Android",
		TimeOut:   time.Minute,
	}

	showType(d)

	q:= impls.Queue{1}
	q.Push(123)
	q.Push(321)
	q.Push(99)

	fmt.Println("isEmpty = ",q.IsEmpty())
	fmt.Println(q)

}
