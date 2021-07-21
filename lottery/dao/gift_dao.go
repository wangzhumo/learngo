package dao

import (
	"com.wangzhumo.lottery/models"
	"xorm.io/xorm"
)

type GiftDao struct {
	engine *xorm.Engine
}

func NewGiftDao(engine *xorm.Engine) *GiftDao {
	return &GiftDao{
		engine: engine,
	}
}

// Get 获取指定ID的Gift信息
func (dao *GiftDao) Get(id int) (*models.LtGift, error) {
	lg := &models.LtGift{Id: id}
	// 查询
	ok, err := dao.engine.Get(lg)
	if !ok && err != nil {
		lg.Id = 0
	}
	return lg, err
}

// GetAll 获取所有的礼品数据
func (dao *GiftDao) GetAll() (list []models.LtGift, err error) {
	lgList := make([]models.LtGift, 0)
	err = dao.engine.
		Asc("sys_status").
		Asc("displayorder").
		Find(&lgList)
	return lgList, err
}

// CountAll 获取所有的礼品数据
func (dao *GiftDao) CountAll() (count int64) {
	count, err := dao.engine.Count(&models.LtGift{})
	if err != nil {
		return 0
	}
	return count
}

// Delete 删除数据 - 实际上是修改状态为 1
func (dao *GiftDao) Delete(id int) (err error) {
	updateData := &models.LtGift{Id: id, SysStatus: 1}
	_, err = dao.engine.ID(updateData.Id).Update(updateData)
	return
}

// Update 强制更新，因为使用Model更新，只会更新不为空的值
func (dao *GiftDao) Update(data *models.LtGift, columes []string) (err error) {
	_, err = dao.engine.ID(data.Id).MustCols(columes...).Update(data)
	return
}

// Insert 写入一个礼物
func (dao *GiftDao) Insert(data *models.LtGift) (err error) {
	_, err = dao.engine.ID(data.Id).Insert(data)
	return
}
