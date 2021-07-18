package metting

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
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

// 用户列表
var userList []string

// 安全 - 加锁 （互斥锁）
var mut sync.Mutex

func Init() {
	app := newApp()
	// init
	userList = []string{}

	err := app.Run(iris.Addr(":8001"))
	if err != nil {
		return
	}

	// 初始化锁
	mut = sync.Mutex{}
}

func (c *LotteryV1Controller) Get() string {
	count := len(userList)
	return fmt.Sprintf("当前人数：%d\n", count)
}

// PostImport 增加抽奖人
func (c *LotteryV1Controller) PostImport() string {
	users := c.Ctx.FormValue("users")
	usersList := strings.Split(users, ",")

	// 并发 - 加锁
	mut.Lock()
	defer mut.Unlock()

	oldCount := len(userList)
	if len(usersList) > 0 {
		for _, v := range usersList {
			user := strings.TrimSpace(v)
			if len(user) > 0 {
				userList = append(userList, user)
			}
		}
		newCount := len(userList)
		return fmt.Sprintf("当前人数：%d ,成功导入用户数：%d\n", newCount, newCount-oldCount)
	}
	return fmt.Sprintf("当前人数：%d\n", oldCount)
}

// GetLottery 抽奖
func (c *LotteryV1Controller) GetLottery() string {
	// 并发 - 加锁
	mut.Lock()
	defer mut.Unlock()

	count := len(userList)
	if count <= 0 {
		return fmt.Sprintf("当前人数：%d 无法开始抽奖", count)
	} else if count == 1 {
		user := userList[0]
		userList = []string{}
		return fmt.Sprintf("当前中奖人：%s ,当前人数：%d\n", user, count-1)
	} else {
		// 简单随机
		seed := time.Now().UnixNano()
		randomIndex := rand.New(rand.NewSource(seed)).Int31n(int32(count))
		user := userList[randomIndex]
		// 修改原始数组
		userList = append(userList[0:randomIndex], userList[randomIndex+1:]...)
		return fmt.Sprintf("当前中奖人：%s ,当前人数：%d\n", user, count-1)
	}
}
