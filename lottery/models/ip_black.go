package models

type IpBlack struct {
	Id         int    `json:"id" xorm:"not null pk autoincr comment('IP黑名单的ID') INT(11)"`
	Ip         string `json:"ip" xorm:"not null comment('ID地址') VARCHAR(255)"`
	Blacktime  int    `json:"blacktime" xorm:"not null comment('黑名单时间') INT(11)"`
	SysUpdate  int    `json:"sys_update" xorm:"comment('更新时间') INT(11)"`
	SysCreated int    `json:"sys_created" xorm:"comment('创建时间') INT(11)"`
}
