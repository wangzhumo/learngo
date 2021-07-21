package dao

import (
	"com.wangzhumo.lottery/models"
	"xorm.io/xorm"
)

type BlackipDao struct {
	engine *xorm.Engine
}

func NewBlackipDao(engine *xorm.Engine) *BlackipDao {
	return &BlackipDao{
		engine: engine,
	}
}

// Get 获取指定ID的Gift信息
func (dao *BlackipDao) Get(id int) (*models.IpBlack, error) {
	lg := &models.IpBlack{Id: id}
	// 查询
	ok, err := dao.engine.Get(lg)
	if !ok && err != nil {
		lg.Id = 0
	}
	return lg, err
}

// GetAll 获取所有的礼品数据
func (dao *BlackipDao) GetAll(id int) (list []models.IpBlack, err error) {
	lgList := make([]models.IpBlack, 0)
	err = dao.engine.
		Asc("sys_status").
		Asc("displayorder").
		Find(&lgList)
	return lgList, err
}

// CountAll 获取所有的礼品数据
func (dao *BlackipDao) CountAll() (count int64) {
	count, err := dao.engine.Count(&models.IpBlack{})
	if err != nil {
		return 0
	}
	return count
}

func (d *BlackipDao) Search(ip string) []models.IpBlack {
	datalist := make([]models.IpBlack, 0)
	err := d.engine.
		Where("ip=?", ip).
		Desc("id").
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

// Update 强制更新，因为使用Model更新，只会更新不为空的值
func (dao *BlackipDao) Update(data *models.IpBlack, columes []string) (err error) {
	_, err = dao.engine.ID(data.Id).MustCols(columes...).Update(data)
	return
}

// Insert 写入一个礼物
func (dao *BlackipDao) Insert(data *models.IpBlack) (err error) {
	_, err = dao.engine.ID(data.Id).Insert(data)
	return
}

// 根据IP获取信息
func (d *BlackipDao) GetByIp(ip string) *models.IpBlack {
	datalist := make([]models.IpBlack, 0)
	err := d.engine.
		Where("ip=?", ip).
		Desc("id").
		Limit(1).
		Find(&datalist)
	if err != nil || len(datalist) < 1 {
		return nil
	} else {
		return &datalist[0]
	}
}