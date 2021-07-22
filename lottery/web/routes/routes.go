package routes

import (
	"com.wangzhumo.lottery/bootstrap"
	"com.wangzhumo.lottery/services"
	"com.wangzhumo.lottery/web/controllers"
	"github.com/kataras/iris/v12/mvc"
)

func Configure(b *bootstrap.Bootstrapper) {
	// 初始化服务
	giftService := services.NewGiftServices()
	blackUserService := services.NewBlackUserService()
	blackIpService := services.NewBlackIpServices()
	codeService := services.NewCodeService()
	resultService := services.NewResultService()
	userDayService := services.NewUserDayService()

	// 路由
	index := mvc.New(b.Party("/"))
	index.Register(userDayService, giftService,
		blackUserService, blackIpService, codeService, resultService)

	index.Handle(new(controllers.IndexController))
}
