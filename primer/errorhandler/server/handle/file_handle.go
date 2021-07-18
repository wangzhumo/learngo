package fileHandle

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefixPath = "/show/"

type UserError string

func (e UserError) Error() string {
	return e.Message()
}

func (e UserError) Message() string {
	return string(e)
}

func FileHttpHandler(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefixPath) != 0 {
		// 如果有这个path,才继续，否则error
		return UserError(fmt.Sprintf("%s has't routing", request.URL.Path))
	}

	// 获取文件path
	path := request.URL.Path[len(prefixPath):]
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
