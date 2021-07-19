package wechat

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// 定义奖品类型
// iota  会自增
const (
	giftTypeCoin = iota
	giftTypeCoupon
	giftTypeCouponFix
	giftTypeRealSamll
	giftTypeRealLarge
)

//Prize 奖品的Item数据
type Prize struct {
	id       int      // 奖品ID
	name     string   //奖品名字
	pic      string   //imageUrl
	link     string   //gift link
	gtype    int      // gift type
	data     string   //奖品的数据（特定的配置信息）
	datalist []string //奖品数据集合（不同优惠卷编码）
	total    int      // gift total number
	left     int      //剩余数量
	inuse    bool     //是否可用
	rate     int      //中奖概率 0 - 9999
	rateMin  int      //大于等于最小中奖编码
	rateMax  int      //小于中奖编码
}

// 互斥锁
var mux *sync.Mutex

// 最大中奖号码
const rateMax = 10000

// log
var logger *log.Logger

func initLog() {
	f, _ := os.Create("./lottery_dome.log")
	logger = log.New(f, "", log.Ldate|log.Lmicroseconds)

	fmt.Println("Logger Init...")
}

// 奖品
var prizeList []*Prize

func initPrize() {
	prizeList = make([]*Prize, 5)
	p1 := Prize{
		id:       1,
		name:     "iPhone",
		pic:      "",
		link:     "",
		gtype:    giftTypeRealLarge,
		data:     "",
		datalist: nil,
		total:    10000,
		left:     10000,
		inuse:    true,
		rate:     10000,
		rateMin:  0,
		rateMax:  0,
	}
	prizeList[0] = &p1

	p2 := Prize{
		id:       2,
		name:     "Android",
		pic:      "",
		link:     "",
		gtype:    giftTypeRealSamll,
		data:     "",
		datalist: nil,
		total:    20,
		left:     20,
		inuse:    true,
		rate:     100,
		rateMin:  0,
		rateMax:  0,
	}
	prizeList[1] = &p2

	p3 := Prize{
		id:       3,
		name:     "200-20",
		pic:      "",
		link:     "",
		gtype:    giftTypeCouponFix,
		data:     "conpon-200-50",
		datalist: nil,
		total:    100,
		left:     100,
		inuse:    true,
		rate:     500,
		rateMin:  0,
		rateMax:  0,
	}
	prizeList[2] = &p3

	p4 := Prize{

		id:       4,
		name:     "50",
		pic:      "",
		link:     "",
		gtype:    giftTypeCoupon,
		data:     "",
		datalist: []string{"conpon1,conpon2,conpon3,conpon4,conpon15"},
		total:    100,
		left:     100,
		inuse:    true,
		rate:     500,
		rateMin:  0,
		rateMax:  0,
	}
	prizeList[3] = &p4

	p5 := Prize{
		id:       5,
		name:     "Coin",
		pic:      "",
		link:     "",
		gtype:    giftTypeCoin,
		data:     "10",
		datalist: nil,
		total:    1000,
		left:     1000,
		inuse:    true,
		rate:     5000,
		rateMin:  0,
		rateMax:  0,
	}
	prizeList[4] = &p5

	rateStart := 0

	// 初始化
	for _, v := range prizeList {
		if !v.inuse {
			continue
		}
		v.rateMin = rateStart
		v.rateMax = rateStart + v.rate
		if v.rateMax >= rateMax {
			v.rateMax = rateMax
			rateStart = 0
		} else {
			rateStart += v.rate
		}
	}

	fmt.Println("Prize Init...")
}

type Lotteryv1Controller struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&Lotteryv1Controller{})
	initLog()
	initPrize()
	mux = &sync.Mutex{}
	return app
}

func InitWechat() {
	app := newApp()
	app.Run(iris.Addr(":8001"))
}

// 返回数量
func (c *Lotteryv1Controller) Get() string {
	count := 0
	total := 0
	for _, v := range prizeList {
		// 可用，而且有库存
		if v.inuse && (v.total == 0 || (v.left > 0 && v.total > 0)) {
			count++
			total += v.total
		}
		fmt.Printf("Prize : id=%d , Name=%s , Min = %d , Max = %d\n", v.id, v.name, v.rateMin, v.rateMax)
	}

	return fmt.Sprintf("当前奖品种类 ：%d ,总数量 ： %d \n", count, total)
}

// 开始抽奖
func (c *Lotteryv1Controller) GetLotterySafe() map[string]interface{} {
	mux.Lock()
	defer mux.Unlock()
	code := 500
	result := make(map[string]interface{})
	result["code"] = code
	// 1.为用户生成code
	randomCode := generateCode()
	// 2.查找这个code对应的奖品
	for _, v := range prizeList {
		// 过滤掉不能使用的prize
		if !v.inuse || (v.total > 0 && v.left <= 0) {
			continue
		}
		if v.rateMin <= int(randomCode) && v.rateMax >= int(randomCode) {
			//区间内，中奖了
			// 3.依据不同的type发奖
			msg := ""
			switch v.gtype {
			case giftTypeCoin:
				code, msg = sendCoin(v)

			case giftTypeCoupon:
				code, msg = sendCompon(v)

			case giftTypeCouponFix:
				code, msg = sendComponFix(v)

			case giftTypeRealSamll:
				code, msg = sendRealSmall(v)

			case giftTypeRealLarge:
				code, msg = sendRealLarge(v)

			}
			// 4.发奖结果判断
			if code == 0 {
				// 成功则生成获奖记录
				savePrizeRecord(code, v.id, v.left, v.name, v.link, msg)
				result["code"] = code
				result["id"] = v.id
				result["name"] = v.name
				result["data"] = msg
				// 然后break
				break
			}
		}

	}
	return result
}

// 开始抽奖
func (c *Lotteryv1Controller) GetLottery() map[string]interface{} {
	code := 500
	result := make(map[string]interface{})
	result["code"] = code
	// 1.为用户生成code
	randomCode := generateCode()
	// 2.查找这个code对应的奖品
	for _, v := range prizeList {
		// 过滤掉不能使用的prize
		if !v.inuse || (v.total > 0 && v.left <= 0) {
			continue
		}
		if v.rateMin <= int(randomCode) && v.rateMax >= int(randomCode) {
			//区间内，中奖了
			// 3.依据不同的type发奖
			msg := ""
			switch v.gtype {
			case giftTypeCoin:
				code, msg = sendCoin(v)

			case giftTypeCoupon:
				code, msg = sendCompon(v)

			case giftTypeCouponFix:
				code, msg = sendComponFix(v)

			case giftTypeRealSamll:
				code, msg = sendRealSmall(v)

			case giftTypeRealLarge:
				code, msg = sendRealLarge(v)

			}
			// 4.发奖结果判断
			if code == 0 {
				// 成功则生成获奖记录
				savePrizeRecord(code, v.id, v.left, v.name, v.link, msg)
				result["code"] = code
				result["id"] = v.id
				result["name"] = v.name
				result["data"] = msg
				// 然后break
				break
			}
		}

	}
	return result
}

func savePrizeRecord(code int, id int, left int, name string, link string, msg string) {
	logger.Printf("Lottery -> code=%d,code=%d,code=%d,code=%s,code=%s,code=%s", code, id, left, name, link, msg)
}

func sendCoin(p *Prize) (code int, msg string) {
	// 0 - 无限量
	if p.total == 0 {
		return 0, p.data
	} else {
		// 否则判断余量
		if p.left > 0 {
			return 0, p.data
		}
	}
	return 500, "奖品已经发完了~"
}

func sendCompon(p *Prize) (code int, msg string) {
	// 0 - 无限量
	if p.total == 0 {
		return 0, p.data
		// 否则判断余量
	} else if p.left > 1 {
		// 减去库存
		left := p.left - 1
		if left > len(p.datalist) {
			return 500, "奖品已经发完了~"
		} else {
			if len(p.datalist) > 1 {
				p.datalist = p.datalist[1:]
			} else {
				p.datalist = []string{}
			}

			p.left = left
			return 0, p.data
		}

	}
	return 500, "奖品已经发完了~"
}

func sendComponFix(p *Prize) (code int, msg string) {
	// 0 - 无限量
	if p.total == 0 {
		return 0, p.data
		// 否则判断余量
	} else if p.left > 1 {
		// 减去库存
		p.left--
		return 0, p.data
	}
	return 500, "奖品已经发完了~"
}

func sendRealLarge(p *Prize) (code int, msg string) {
	// 0 - 无限量
	if p.total == 0 {
		return 0, p.data
		// 否则判断余量
	} else if p.left > 1 {
		// 减去库存
		p.left--
		return 0, p.data
	}
	return 500, "奖品已经发完了~"
}

func sendRealSmall(p *Prize) (code int, msg string) {
	// 0 - 无限量
	if p.total == 0 {
		return 0, p.data
		// 否则判断余量
	} else if p.left > 1 {
		// 减去库存
		p.left--
		return 0, p.data
	}
	return 500, "奖品已经发完了~"
}

// 每一个人的code - 和抽奖的code匹配
func generateCode() int32 {
	seed := time.Now().UnixNano()
	i := rand.New(rand.NewSource(seed)).Int31n(int32(rateMax))
	return i
}
