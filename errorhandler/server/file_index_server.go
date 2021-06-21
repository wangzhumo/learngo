package fileServer

import (
	"log"
	"net/http"
	"os"
	fileHandle "wangzhumo.com/learngo/errorhandler/server/handle"
)

// 定义一个函数返回为处理http的函数
type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 处理错误
func errHandler(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {
		// recover,处理一些意外情况
		defer func() {
			if r := recover(); r != nil {
				// 打印出去
				log.Printf("File Server Error : %v ", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		// 处理所有的错误
		code := http.StatusOK
		if err != nil {
			// 如果是userError 我们可以先处理掉
			if userErr, ok := err.(UserError); ok {
				// 那么直接输出给用户即可
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			// log
			log.Printf("Error handler : %s", err.Error())
			switch {
			case os.IsNotExist(err):
				// 如果不存在
				code = http.StatusNotFound
			case os.IsPermission(err):
				// 如果没有权限
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}

}

type UserError interface {
	error
	Message() string
}

// RunFileServer 显示文件
func RunFileServer() {
	http.HandleFunc("/", errHandler(fileHandle.FileHttpHandler))

	// http://localhost:9999/show/errorhandler/defer/fib.txt
	http.ListenAndServe(":9999", nil)
}
