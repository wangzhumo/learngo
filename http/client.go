package httpClient

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const httpUrl = "https://www.imooc.com"

func httpClient() {
	// request build
	request, err := http.NewRequest(
		http.MethodGet,
		httpUrl,
		nil,
	)
	if err != nil {
		panic(err)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T)")

	// client
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req)
			return nil
		},
	}
	resp, err := client.Do(request)

	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		panic(err)
	}

	//dump to file
	file, err := os.OpenFile("http/index.html", os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	witer := bufio.NewWriter(file)
	defer witer.Flush()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	witer.Write(bodyBytes)
}

func RunHttpClient() {
	httpClient()
}
