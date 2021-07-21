package models

type UserBlack struct {
	Id         int    `json:"id" xorm:"not null pk autoincr comment('用户ID') INT(11)"`
	Username   string `json:"username" xorm:"not null comment('用户Nick') VARCHAR(255)"`
	Blacktime  int    `json:"blacktime" xorm:"not null comment('黑名单时间') INT(11)"`
	Realname   string `json:"realname" xorm:"comment('联系人') VARCHAR(50)"`
	Mobile     string `json:"mobile" xorm:"comment('用户手机号') VARCHAR(50)"`
	Address    string `json:"address" xorm:"comment('地址') VARCHAR(255)"`
	SysUpdate  int    `json:"sys_update" xorm:"comment('更新时间') INT(11)"`
	SysCreated int    `json:"sys_created" xorm:"comment('创建时间') INT(11)"`
	SysIp      string `json:"sys_ip" xorm:"comment('用户的IP') VARCHAR(50)"`
}
