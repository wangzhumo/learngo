package dao

import (
	"com.wangzhumo.lottery/models"
	"xorm.io/xorm"
)

type ResultDao struct {
	engine *xorm.Engine
}

func NewResultDao(engine *xorm.Engine) *ResultDao {
	return &ResultDao{
		engine: engine,
	}
}

func (d *ResultDao) Get(id int) *models.GiftRecord {
	data := &models.GiftRecord{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *ResultDao) GetAll(page, size int) []models.GiftRecord {
	offset := (page - 1) * size
	datalist := make([]models.GiftRecord, 0)
	err := d.engine.
		Desc("id").
		Limit(size, offset).
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *ResultDao) CountAll() int64 {
	num, err := d.engine.
		Count(&models.GiftRecord{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *ResultDao) GetNewPrize(size int, giftIds []int) []models.GiftRecord {
	datalist := make([]models.GiftRecord, 0)
	err := d.engine.
		In("gift_id", giftIds).
		Desc("id").
		Limit(size).
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *ResultDao) SearchByGift(giftId, page, size int) []models.GiftRecord {
	offset := (page - 1) * size
	datalist := make([]models.GiftRecord, 0)
	err := d.engine.
		Where("gift_id=?", giftId).
		Desc("id").
		Limit(size, offset).
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *ResultDao) SearchByUser(uid, page, size int) []models.GiftRecord {
	offset := (page - 1) * size
	datalist := make([]models.GiftRecord, 0)
	err := d.engine.
		Where("uid=?", uid).
		Desc("id").
		Limit(size, offset).
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *ResultDao) CountByGift(giftId int) int64 {
	num, err := d.engine.
		Where("gift_id=?", giftId).
		Count(&models.GiftRecord{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *ResultDao) CountByUser(uid int) int64 {
	num, err := d.engine.
		Where("uid=?", uid).
		Count(&models.GiftRecord{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *ResultDao) Delete(id int) error {
	data := &models.GiftRecord{Id: id, SysStatus: 1}
	_, err := d.engine.ID(data.Id).Update(data)
	return err
}

func (d *ResultDao) Update(data *models.GiftRecord, columns []string) error {
	_, err := d.engine.ID(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *ResultDao) Insert(data *models.GiftRecord) error {
	_, err := d.engine.Insert(data)
	return err
}
