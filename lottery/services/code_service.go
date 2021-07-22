package services

import (
	"com.wangzhumo.lottery/dao"
	"com.wangzhumo.lottery/models"
)

type CodeService interface {
	Get(id int) *models.LtCoupon
	GetAll(page, size int) []models.LtCoupon
	CountAll() int64
	CountByGift(giftId int) int64
	Search(giftId int) []models.LtCoupon
	Delete(id int) error
	Update(user *models.LtCoupon, columns []string) error
	Insert(user *models.LtCoupon) error
	NextUsingCode(giftId, codeId int) *models.LtCoupon
	UpdateByCode(data *models.LtCoupon, columns []string) error
}

type codeService struct {
	dao *dao.LtCouponDao
}

func (c *codeService) Get(id int) (*models.LtCoupon, error) {
	return c.dao.Get(id)
}

func (c *codeService) GetAll(page, size int) ([]models.LtCoupon, error) {
	return c.dao.GetAll(page, size)
}

func (c *codeService) CountAll() int64 {
	return c.dao.CountAll()
}

func (c *codeService) CountByGift(giftId int) int64 {
	return c.dao.CountByGift(giftId)
}

func (c *codeService) Search(giftId int) []models.LtCoupon {
	return c.dao.Search(giftId)
}

func (c *codeService) Delete(id int) error {
	return c.dao.Delete(id)
}

func (c *codeService) Update(code *models.LtCoupon, columns []string) error {
	return c.dao.Update(code, columns)
}

func (c *codeService) Insert(code *models.LtCoupon) error {
	return c.dao.Insert(code)
}

func (c *codeService) NextUsingCode(giftId, codeId int) *models.LtCoupon {
	return c.dao.NextUsingCode(giftId, codeId)
}

func (c *codeService) UpdateByCode(data *models.LtCoupon, columns []string) error {
	return c.dao.UpdateByCode(data, columns)
}
