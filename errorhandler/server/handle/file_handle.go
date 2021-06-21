package fileHandle

import (
	"io/ioutil"
	"net/http"
	"os"
)

func FileHttpHandler(writer http.ResponseWriter, request *http.Request) error {
	// 获取文件path
	path := request.URL.Path[len("/show/"):]
	// 打开文件
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return err
	}

	// 获取文件内容
	cont, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	// 写入Response
	writer.Write(cont)
	return nil
}
