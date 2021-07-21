package dao

import (
	"com.wangzhumo.lottery/models"
	"xorm.io/xorm"
)

type CountDao struct {
	engine *xorm.Engine
}

func NewCountDao(engine *xorm.Engine) *CountDao {
	return &CountDao{
		engine: engine,
	}
}

// Get 获取指定ID的Gift信息
func (dao *CountDao) Get(id int) (*models.CountRecord, error) {
	lg := &models.CountRecord{Id: id}
	// 查询
	ok, err := dao.engine.Get(lg)
	if !ok && err != nil {
		lg.Id = 0
	}
	return lg, err
}

// GetAll 获取所有的Day数据
func (dao *CountDao) GetAll(page int,size int) (list []models.CountRecord, err error) {
	offset:= (page-1)*size
	lgList := make([]models.CountRecord, 0)
	err = dao.engine.
		Desc("id").
		Limit(size,offset).
		Find(&lgList)
	return lgList, err
}

// Search 查询某个用户
func (dao *CountDao) Search(uid, day int) []models.CountRecord {
	datalist := make([]models.CountRecord, 0)
	err := dao.engine.
		Where("uid=?", uid).
		Where("day=?", day).
		Desc("id").
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}


// CountAll 获取所有数据的总数
func (dao *CountDao) CountAll() (count int64) {
	count, err := dao.engine.Count(&models.CountRecord{})
	if err != nil {
		return 0
	}
	return count
}

//// Delete 删除数据 - 实际上是修改状态为 1
//func (dao *CountDao) Delete(id int) (err error) {
//	updateData := &models.CountRecord{Id: id, SysStatus: 1}
//	_, err = dao.engine.ID(updateData.Id).Update(updateData)
//	return
//}

// Update 强制更新，因为使用Model更新，只会更新不为空的值
func (dao *CountDao) Update(data *models.CountRecord, columes []string) (err error) {
	_, err = dao.engine.ID(data.Id).MustCols(columes...).Update(data)
	return
}

// Insert 写入一个礼物
func (dao *CountDao) Insert(data *models.CountRecord) (err error) {
	_, err = dao.engine.ID(data.Id).Insert(data)
	return
}
