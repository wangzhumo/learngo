package models

type GiftRecord struct {
	Id         int    `json:"id" xorm:"not null pk autoincr comment('其他奖品的ID，主键自增') INT(11)"`
	GiftId     int    `json:"gift_id" xorm:"not null comment('奖品ID - 关联gift表中的ID') INT(11)"`
	GiftName   string `json:"gift_name" xorm:"not null comment('奖品名称') VARCHAR(255)"`
	GiftType   int    `json:"gift_type" xorm:"not null comment('奖品Type - 关联gift表中的gtype') INT(11)"`
	Uid        int    `json:"uid" xorm:"not null comment('用户ID - 发放给谁了') INT(11)"`
	Username   string `json:"username" xorm:"not null comment('用户Nick') VARCHAR(255)"`
	PrizeCode  int    `json:"prize_code" xorm:"not null comment('用户抽奖时的code码') INT(11)"`
	GiftData   string `json:"gift_data" xorm:"comment('获奖说明') VARCHAR(255)"`
	SysStatus  int    `json:"sys_status" xorm:"not null comment('状态 0正常 1删除 2异常') TINYINT(4)"`
	SysCreated int    `json:"sys_created" xorm:"comment('创建时间') INT(11)"`
	SysIp      string `json:"sys_ip" xorm:"comment('用户的IP') VARCHAR(50)"`
}
