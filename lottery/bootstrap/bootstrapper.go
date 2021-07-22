package bootstrap

import (
	"com.wangzhumo.lottery/conf"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"time"
)

const (
	StaticAssets = "./public/"
	StaticRoute = "public"
	Favicon      = "favicon.ico"
)

type Configurator func(*Bootstrapper)

// Bootstrapper 封装
// 使用Go内建的嵌入机制(匿名嵌入)，允许类型之前共享代码和数据
type Bootstrapper struct {
	*iris.Application           // 继承
	AppName           string    // 名字
	AppOwner          string    // 负责人
	AppSpawnDate       time.Time // 更新时间
}

// New 创建一个Iris的实例 - Bootstrapper
func New(appName string, appOwner string, cfgs ...Configurator) *Bootstrapper {
	// 创建实例
	bootstrapper := &Bootstrapper{
		Application: iris.New(),
		AppName:     appName,
		AppOwner:    appOwner,
		AppSpawnDate: time.Time{},
	}
	// 调用其他人的初始化方法
	for _, configurator := range cfgs {
		configurator(bootstrapper)
	}
	// 返回即可
	return bootstrapper
}

// Configure 提供给外部修改配置的方法
func (bootstrapper *Bootstrapper) Configure(configs ...Configurator) {
	for _, configurator := range configs {
		configurator(bootstrapper)
	}
}

// Bootstrap 启动应用程序必须的组件
func (bootstrapper *Bootstrapper) Bootstrap() *Bootstrapper {
	// 绑定View
	bootstrapper.SetupViews("./views")
	// 开启默认错误处理
	bootstrapper.SetupErrorHandler()
	// 设置Favicon
	bootstrapper.Favicon(StaticAssets + Favicon)
	// 启动静态服务
	bootstrapper.HandleDir(StaticRoute,StaticAssets)
	// 定时服务
	bootstrapper.SetupCron()
	// Recover服务
	bootstrapper.Use(recover.New())
	bootstrapper.Use(logger.New())
	return bootstrapper
}

// Listen 启动监听
func (bootstrapper *Bootstrapper) Listen(addr string,configurators ...iris.Configurator)  {
	bootstrapper.Run(iris.Addr(addr),configurators...)
}


// SetupCron 定时任务
func (bootstrapper *Bootstrapper) SetupCron() {

}

// SetupViews 初始化模版
func (bootstrapper *Bootstrapper) SetupViews(viewDir string) {
	// 设置参数，获取Engine
	htmlEngine := iris.HTML(viewDir, ".html").
		Layout("shared/layout.html")
	// 热加载
	htmlEngine.Reload(true)
	// 通用的方法
	htmlEngine.AddFunc("FromUnixTimeShort", func(timestamp int) string {
		dt := time.Unix(int64(timestamp), int64(0))
		return dt.Format(conf.SysTimeFormatShort)
	})
	htmlEngine.AddFunc("FromUnixTime", func(timestamp int) string {
		dt := time.Unix(int64(timestamp), int64(0))
		return dt.Format(conf.SysTimeFormat)
	})
	// 注入
	bootstrapper.RegisterView(htmlEngine)
}

// SetupErrorHandler 一些通用的异常处理
func (bootstrapper *Bootstrapper) SetupErrorHandler() {
	bootstrapper.OnAnyErrorCode(func(ctx iris.Context) {
		errorData := iris.Map{
			"app":     bootstrapper.AppName,
			"status":  ctx.GetStatusCode(),
			"message": ctx.Values().GetString("message"),
		}
		if jsonOutput := ctx.URLParamExists("json"); jsonOutput {
			_, _ = ctx.JSON(errorData)
			return
		}
		ctx.ViewData("Err", errorData)
		ctx.ViewData("Title", "Error")
		ctx.View("shared/error.html")
	})
}
