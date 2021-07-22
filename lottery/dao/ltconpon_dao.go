package dao

import (
	"com.wangzhumo.lottery/models"
	"xorm.io/xorm"
)

type LtCouponDao struct {
	engine *xorm.Engine
}

func NewLtCouponDao(engine *xorm.Engine) *LtCouponDao {
	return &LtCouponDao{
		engine: engine,
	}
}

// Get 获取指定ID的Gift信息
func (dao *LtCouponDao) Get(id int) (*models.LtCoupon, error) {
	lg := &models.LtCoupon{Id: id}
	// 查询
	ok, err := dao.engine.Get(lg)
	if !ok && err != nil {
		lg.Id = 0
	}
	return lg, err
}

// GetAll 获取所有的礼品数据
func (dao *LtCouponDao) GetAll(page int, size int) (list []models.LtCoupon, err error) {
	offset := (page - 1) * size
	lgList := make([]models.LtCoupon, 0)
	err = dao.engine.
		Desc("id").
		Limit(size, offset).
		Find(&lgList)
	return lgList, err
}

// CountAll 获取所有的礼品数据
func (dao *LtCouponDao) CountAll() (count int64) {
	count, err := dao.engine.Count(&models.LtCoupon{})
	if err != nil {
		return 0
	}
	return count
}

func (dao *LtCouponDao) CountByGift(giftId int) int64 {
	num, err := dao.engine.
		Where("gift_id=?", giftId).
		Count(&models.LtCoupon{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (dao *LtCouponDao) Search(giftId int) []models.LtCoupon {
	datalist := make([]models.LtCoupon, 0)
	err := dao.engine.
		Where("gift_id=?", giftId).
		Desc("id").
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

// Delete 删除数据 - 实际上是修改状态为 1
func (dao *LtCouponDao) Delete(id int) (err error) {
	updateData := &models.LtCoupon{Id: id, SysStatus: 1}
	_, err = dao.engine.ID(updateData.Id).Update(updateData)
	return
}

// Update 强制更新，因为使用Model更新，只会更新不为空的值
func (dao *LtCouponDao) Update(data *models.LtCoupon, columes []string) (err error) {
	_, err = dao.engine.ID(data.Id).MustCols(columes...).Update(data)
	return
}

// Insert 写入一个Coupon礼物
func (dao *LtCouponDao) Insert(data *models.LtCoupon) (err error) {
	_, err = dao.engine.ID(data.Id).Insert(data)
	return
}

// NextUsingCode 找到下一个可用的最小的优惠券
func (dao *LtCouponDao) NextUsingCode(giftId, codeId int) *models.LtCoupon {
	datalist := make([]models.LtCoupon, 0)
	err := dao.engine.Where("gift_id=?", giftId).
		Where("sys_status=?", 0).
		Where("id>?", codeId).
		Asc("id").Limit(1).
		Find(&datalist)
	if err != nil || len(datalist) < 1 {
		return nil
	} else {
		return &datalist[0]
	}
}

//UpdateByCode 根据唯一的code来更新
func (dao *LtCouponDao) UpdateByCode(data *models.LtCoupon, columns []string) error {
	_, err := dao.engine.Where("code=?", data.Code).
		MustCols(columns...).Update(data)
	return err
}
