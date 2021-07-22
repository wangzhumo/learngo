package controllers

import (
	"com.wangzhumo.lottery/models"
	"com.wangzhumo.lottery/services"
	"github.com/kataras/iris/v12"
)

type IndexController struct {
	Ctx            iris.Context
	ServiceUser    services.UserService
	ServiceIp      services.BlackIpService
	GiftServices   services.GiftServices
	CodeService    services.CodeService
	UserDayService services.UserDayService
	ResultService  services.ResultService
}

// Get 首页
func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "welcome to Lottery system,<a href='/public/index.html'>开始抽奖</a>"
}

// GifGiftList 获取礼物列表
func (c *IndexController) GifGiftList() map[string]interface{} {
	response := make(map[string]interface{})
	response["code"] = 0
	response["message"] = ""
	// find gift list
	gifts, err := c.GiftServices.GetAll()
	giftList := make([]models.LtGift, 0)
	if err == nil {
		for _, gift := range gifts {
			// 写入返回值
			if gift.SysStatus == 0 {
				giftList = append(giftList, gift)
			}
		}
	}
	response["data"] = giftList
	return response
}


// GifNewGiftList 获取最近中奖
func (c *IndexController) GifNewGiftList() map[string]interface{} {
	response := make(map[string]interface{})
	response["code"] = 0
	response["message"] = ""
	// find gift list
	gifts, err := c.GiftServices.GetAll()
	giftList := make([]models.LtGift, 0)
	if err == nil {
		for _, gift := range gifts {
			// 写入返回值
			if gift.SysStatus == 0 {
				giftList = append(giftList, gift)
			}
		}
	}
	response["data"] = giftList
	return response
}
