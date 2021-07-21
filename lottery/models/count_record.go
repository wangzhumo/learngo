package models

type CountRecord struct {
	Id         int `json:"id" xorm:"not null pk autoincr comment('每日抽奖次数表的ID') INT(11)"`
	Uid        int `json:"uid" xorm:"not null comment('用户ID') INT(11)"`
	Day        int `json:"day" xorm:"not null comment('日期') INT(11)"`
	Num        int `json:"num" xorm:"comment('抽奖次数') INT(11)"`
	SysCreated int `json:"sys_created" xorm:"comment('创建时间') INT(11)"`
	SysUpdate  int `json:"sys_update" xorm:"comment('更新时间') INT(11)"`
}
