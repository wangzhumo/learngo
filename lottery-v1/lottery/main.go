package lottery

import (
	"fmt"
	"math/rand"
	"strings"
	"sync/atomic"
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
	// init
	err := app.Run(iris.Addr(":8001"))
	if err != nil {
		return
	}
}

// 奖品的中奖概率
type Prate struct {
	Rate  int    // 概率  0-10000
	Total int    //奖品数量
	Code  int    //开始位置（含）
	Code1 int    //终止位置（含）
	Left  *int32 //剩余
}

// 奖品
var prizeList []string = []string{
	"一等奖，Apple全家桶",
	"二等奖，Microsoft全家桶",
	"三等奖，Huawei Mate 50",
	"",
}

var left1 int32 = 10000
var left2 int32 = 2
var left3 int32 = 10
var left4 int32 = 0

// 中将概率设置
var rateList []Prate = []Prate{
	Prate{Rate: 10000, Total: 10000, Code: 0, Code1: 10000, Left: &left1},
	Prate{Rate: 2, Total: 2, Code: 1, Code1: 2, Left: &left2},
	Prate{Rate: 5, Total: 10, Code: 3, Code1: 5, Left: &left3},
	Prate{Rate: 100, Total: 0, Code: 0, Code1: 9999, Left: &left4},
}

func (c *LotteryV1Controller) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return fmt.Sprintf("奖品列表：<br/> %s", strings.Join(prizeList, "<br/>\n"))
}

func (c *LotteryV1Controller) GetDebug() string {
	return fmt.Sprintf("奖品概率：%v\n", rateList)
}

func (c *LotteryV1Controller) GetLottery() string {
	//1.确定code
	userCode := generateCode()
	//2.比对
	var prize string
	var prizeRate *Prate
	for i, v := range prizeList {
		//2.1.取出奖品概率
		rate := &rateList[i]
		if rate.Code <= int(userCode) || rate.Code1 >= int(userCode) {
			//2.2确定中奖，继续发奖
			prize = v
			prizeRate = rate
			break
		}
	}

	// 3.发奖
	if prize == "" {
		prize = "很遗憾，请再试一次吧~"
		return prize
	}
	// 否则继续发奖
	if prizeRate.Total == 0 {
		return prize
	} else if *prizeRate.Left > 0 {
		i := atomic.AddInt32(prizeRate.Left, -1)
		if i > 0 {

			return prize
		}

	}
	return "很遗憾，请再试一次吧~"
}

// 每一个人的code - 和抽奖的code匹配
func generateCode() int32 {
	seed := time.Now().UnixNano()
	i := rand.New(rand.NewSource(seed)).Int31n(int32(10000))
	return i
}
