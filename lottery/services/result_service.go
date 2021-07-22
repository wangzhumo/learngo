package services

import (
	"com.wangzhumo.lottery/dao"
	"com.wangzhumo.lottery/datasource"
	"com.wangzhumo.lottery/models"
)

type ResultService interface {
	Get(id int) *models.GiftRecord

	GetAll(page, size int) []models.GiftRecord

	CountAll() int64

	GetNewPrize(size int, giftIds []int) []models.GiftRecord

	SearchByGift(giftId, page, size int) []models.GiftRecord

	SearchByUser(uid, page, size int) []models.GiftRecord

	CountByGift(giftId int) int64

	CountByUser(uid int) int64

	Delete(id int) error

	Update(data *models.GiftRecord, columns []string) error

	Insert(data *models.GiftRecord) error
}

type resultService struct {
	dao *dao.ResultDao
}

func (r resultService) Get(id int) *models.GiftRecord {
	return r.dao.Get(id)
}

func (r resultService) GetAll(page, size int) []models.GiftRecord {
	return r.dao.GetAll(page, size)
}

func (r resultService) CountAll() int64 {
	return r.dao.CountAll()
}

func (r resultService) GetNewPrize(size int, giftIds []int) []models.GiftRecord {
	return r.dao.GetNewPrize(size,giftIds)
}

func (r resultService) SearchByGift(giftId, page, size int) []models.GiftRecord {
	return r.dao.SearchByGift(giftId,page,size)
}

func (r resultService) SearchByUser(uid, page, size int) []models.GiftRecord {
	return r.dao.SearchByUser(uid,page,size)
}

func (r resultService) CountByGift(giftId int) int64 {
	return r.dao.CountByGift(giftId)
}

func (r resultService) CountByUser(uid int) int64 {
	return r.dao.CountByUser(uid)
}

func (r resultService) Delete(id int) error {
	return r.dao.Delete(id)
}

func (r resultService) Update(data *models.GiftRecord, columns []string) error {
	return r.dao.Update(data,columns)
}

func (r resultService) Insert(data *models.GiftRecord) error {
	return r.dao.Insert(data)
}


func NewResultService() ResultService {
	return &resultService{dao: dao.NewResultDao(datasource.InstanceDB())}
}
