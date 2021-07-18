package crawler

import (
	"bytes"
	"fmt"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
)

const startPointUrl = "https://www.zhenai.com/zhenghun"

// RunCrawlerTask 开启爬虫任务
func RunCrawlerTask() {
	// 读取空间信息
	resp, err := http.Get(startPointUrl)
	if err != nil {
		panic(err)
	}

	// 记得关闭
	if resp != nil {
		defer resp.Body.Close()
	}

	// 判读类型
	if resp.StatusCode != http.StatusOK {
		fmt.Println("error : status code ", resp.StatusCode)
		return
	}

	// 获取编码类型

	e := deterEncoding(resp.Body)
	uft8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	// 读取内容
	b, err := ioutil.ReadAll(uft8Reader)
	if err != nil {
		panic(err)
	}

	// 打印
	fmt.Printf("%s \n", b)
}

// 获取编码类型
func deterEncoding(r io.Reader) encoding.Encoding {
	// io.Reader被读取的部分不能再次读取，所以需要复制
	var buf bytes.Buffer
	io.TeeReader(r, &buf)
	all, _ := ioutil.ReadAll(&buf)
	e, name, _ := charset.DetermineEncoding(all, "")
	fmt.Println("Encoding Name : ", name)
	return e
}
