package main

import (
	"com.wangzhumo.lottery/bootstrap"
	"com.wangzhumo.lottery/web/middleware/identity"
	"com.wangzhumo.lottery/web/routes"
	"fmt"
)

const Port = 8001

func newApp() *bootstrap.Bootstrapper {
	// 初始化Application
	app:= bootstrap.New("Lottery System", "wangzhumo")
	app.Bootstrap()

	// 中间件
	app.Configure(identity.Configure)
	// 自己的Route
	app.Configure(routes.Configure)
	return app
}


func main() {
	app:= newApp()
	app.Listen(fmt.Sprintf(":%d",Port))
}
