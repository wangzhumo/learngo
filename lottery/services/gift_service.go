package services

import (
	"com.wangzhumo.lottery/dao"
	"com.wangzhumo.lottery/models"
)

// GiftServices 使用接口定义了Gift可以对外提供的功能
// 1.封装利于替换
// 2.可以有多种实现
type GiftServices interface {
	Get(id int) (*models.LtGift, error)
	GetAll() (list []models.LtGift, err error)
	CountAll() int64
	Delete(id int) error
	Update(data *models.LtGift, columns []string) error
	Insert(data *models.LtGift) error
}

// 具体的实现
type giftServices struct {
	dao *dao.GiftDao
}

func (g *giftServices) Get(id int) (*models.LtGift, error) {
	return g.dao.Get(id)
}

func (g *giftServices) GetAll() (list []models.LtGift, err error) {
	return g.dao.GetAll()
}

func (g *giftServices) CountAll() int64 {
	return g.dao.CountAll()
}

func (g *giftServices) Delete(id int) error {
	return g.dao.Delete(id)
}

func (g *giftServices) Update(data *models.LtGift, columns []string) error {
	return g.dao.Update(data, columns)
}

func (g *giftServices) Insert(data *models.LtGift) error {
	return g.dao.Insert(data)
}

func NewGiftServices() GiftServices {
	return &giftServices{
		dao: dao.NewGiftDao(nil),
	}
}
