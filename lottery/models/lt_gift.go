package models

type LtGift struct {
	Id           int    `json:"id" xorm:"not null pk autoincr comment('奖品ID，主键自增') INT(11)"`
	Title        string `json:"title" xorm:"not null comment('奖品名称') VARCHAR(255)"`
	PrizeNum     int    `json:"prize_num" xorm:"not null default -1 comment('奖品数量  0无限个 >0 限量 <0无奖品') INT(11)"`
	PrizeCode    string `json:"prize_code" xorm:"not null comment('奖品的中奖概率 0-9999->100%  0-0 万分之一') VARCHAR(50)"`
	PrizeTime    int    `json:"prize_time" xorm:"comment('发奖周期，单位：天') INT(11)"`
	Img          string `json:"img" xorm:"comment('礼物图片') VARCHAR(255)"`
	DisplayOrder int    `json:"display_order" xorm:"comment('奖品的排序') INT(11)"`
	Gtype        int    `json:"gtype" xorm:"comment('礼物类型 0虚拟币 1券 2实物 3大奖') INT(11)"`
	Gdata        string `json:"gdata" xorm:"comment('扩展数据：如虚拟币数量') VARCHAR(255)"`
	TimeBegin    int    `json:"time_begin" xorm:"comment('可用时间-开始') INT(11)"`
	TimeEnd      int    `json:"time_end" xorm:"comment('结束时间') INT(11)"`
	PrizeData    string `json:"prize_data" xorm:"comment('发奖计划（json数据）') MEDIUMTEXT"`
	PrizeBegin   int    `json:"prize_begin" xorm:"comment('发奖周期-开始') INT(11)"`
	PrizeEnd     int    `json:"prize_end" xorm:"comment('发奖周期-结束') INT(11)"`
	SysStatus    int    `json:"sys_status" xorm:"not null comment('状态 0正常 1移除') TINYINT(4)"`
	SysCreated   int    `json:"sys_created" xorm:"comment('创建时间') INT(11)"`
	SysUpdate    int    `json:"sys_update" xorm:"comment('更新时间') INT(11)"`
	SysIp        string `json:"sys_ip" xorm:"comment('操作人IP') VARCHAR(50)"`
}
