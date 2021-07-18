package ticket

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// LotteryV1Controller 注入
type LotteryV1Controller struct {
	Ctx iris.Context
}

// 开启Iris - add LotteryV1Controller
func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&LotteryV1Controller{})
	return app
}

func Init() {
	app := newApp()
	err := app.Run(iris.Addr(":8001"))
	if err != nil {
		return
	}
}

// 提前公布中奖号 - 刮开对比
func (c *LotteryV1Controller) GetOpen() string {
	c.Ctx.Header("Content-Type", "text/html")
	seed := time.Now().UnixNano()                   // rand内部运算的随机数
	code := rand.New(rand.NewSource(seed)).Intn(10) // rand计算得到的随机数
	var prize string
	switch {
	case code == 1:
		prize = "一等奖"
	case code >= 2 && code <= 3:
		prize = "二等奖"
	case code >= 4 && code <= 6:
		prize = "三等奖"
	default:
		return fmt.Sprintf("尾号为1获得一等奖<br/>"+
			"尾号为2或者3获得二等奖<br/>"+
			"尾号为4/5/6获得三等奖<br/>"+
			"code=%d<br/>"+
			"很遗憾，没有获奖", code)
	}
	return fmt.Sprintf("尾号为1获得一等奖<br/>"+
		"尾号2或者3获得二等奖<br/>"+
		"尾号4/5/6获得三等奖<br/>"+
		"code=%d<br/>"+
		"恭喜你获得:%s", code, prize)
}

// 提前选定号码 - 后续抽取
func (c *LotteryV1Controller) GetPrize() string {
	c.Ctx.Header("Content-Type", "text/html")
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	var prize [7]int
	// 红色球，1-33
	for i := 0; i < 6; i++ {
		prize[i] = r.Intn(33) + 1
	}
	// 最后一位的蓝色球，1-16
	prize[6] = r.Intn(16) + 1
	return fmt.Sprintf("今日开奖号码是： %v", prize)
}
