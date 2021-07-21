package models

type LtCoupon struct {
	Id         int    `json:"id" xorm:"not null pk autoincr comment('其他奖品的ID，主键自增') INT(11)"`
	GiftId     int    `json:"gift_id" xorm:"not null default -1 comment('奖品ID - 关联lt_gift表中的ID') INT(11)"`
	Code       string `json:"code" xorm:"not null comment('奖品的编码号') VARCHAR(255)"`
	SysStatus  int    `json:"sys_status" xorm:"not null comment('状态 0正常 1移除 2已发放') TINYINT(4)"`
	SysCreated int    `json:"sys_created" xorm:"comment('创建时间') INT(11)"`
	SysUpdate  int    `json:"sys_update" xorm:"comment('更新时间') INT(11)"`
}
