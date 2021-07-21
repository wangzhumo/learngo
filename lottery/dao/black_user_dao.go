package dao

import (
	"com.wangzhumo.lottery/models"
	"xorm.io/xorm"
)

type BlackUserDao struct {
	engine *xorm.Engine
}

func NewBlackUserDao(engine *xorm.Engine) *BlackUserDao {
	return &BlackUserDao{
		engine: engine,
	}
}

// Get 获取指定ID的Gift信息
func (dao *BlackUserDao) Get(id int) (*models.UserBlack, error) {
	lg := &models.UserBlack{Id: id}
	// 查询
	ok, err := dao.engine.Get(lg)
	if !ok && err != nil {
		lg.Id = 0
	}
	return lg, err
}

// GetAll 获取所有的礼品数据
func (dao *BlackUserDao) GetAll(id int) (list []models.UserBlack, err error) {
	lgList := make([]models.UserBlack, 0)
	err = dao.engine.
		Asc("sys_status").
		Asc("displayorder").
		Find(&lgList)
	return lgList, err
}

// CountAll 获取所有的礼品数据
func (dao *BlackUserDao) CountAll() (count int64) {
	count, err := dao.engine.Count(&models.UserBlack{})
	if err != nil {
		return 0
	}
	return count
}



// Update 强制更新，因为使用Model更新，只会更新不为空的值
func (dao *BlackUserDao) Update(data *models.UserBlack, columes []string) (err error) {
	_, err = dao.engine.ID(data.Id).MustCols(columes...).Update(data)
	return
}

// Insert 写入一个礼物
func (dao *BlackUserDao) Insert(data *models.UserBlack) (err error) {
	_, err = dao.engine.ID(data.Id).Insert(data)
	return
}
