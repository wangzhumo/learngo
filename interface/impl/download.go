package impls

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Download struct {
	DownloadUrl string
	UserAgent   string
	TimeOut     time.Duration
}

func (download *Download) Get(url string) string {
	// start http
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	// read response
	result, err := httputil.DumpResponse(response, true)

	response.Body.Close()

	if err != nil {
		panic(err)
	}

	return string(result)
}
