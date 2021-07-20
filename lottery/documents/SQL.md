# 数据库表设计

## 奖品

```sql

create table prize
(
	id int auto_increment comment '奖品ID，主键自增',
	title varchar(255) not null comment '奖品名称',
	prize_num int default -1 not null comment '奖品数量  0无限个 >0 限量 <0无奖品',
	prize_code varchar(50) not null comment '奖品的中奖概率 0-9999->100%  0-0 万分之一',
	prize_time int null comment '发奖周期，单位：天',
	img varchar(255) null comment '礼物图片',
	display_order int null comment '奖品的排序',
	gtype int null comment '礼物类型 0虚拟币 1券 2实物 3大奖',
	gdata varchar(255) null comment '扩展数据：如虚拟币数量',
	time_begin int null comment '可用时间-开始',
	time_end int null comment '结束时间',
	prize_data mediumtext null comment '发奖计划（json数据）',
	prize_begin int null comment '发奖周期-开始',
	prize_end int null comment '发奖周期-结束',
	sys_status tinyint not null comment '状态 0正常 1移除',
	sys_create int null comment '创建时间',
	sys_update int null comment '更新时间',
	sys_ip varchar(50) null comment '操作人IP',
	constraint prize_pk
		primary key (id)
) CHARACTER SET utf8 COLLATE utf8_general_ci comment '奖品数据表';

```


## 其他奖品

```sql

create table coupon
(
	id int auto_increment comment '其他奖品的ID，主键自增',
	gift_id int default -1 not null comment '奖品ID - 关联gift表中的ID',
	code varchar(255) not null comment '奖品的编码号',
	sys_status tinyint not null comment '状态 0正常 1移除 2已发放',
	sys_created int null comment '创建时间',
	sys_update int null comment '更新时间',
	constraint prize_pk
		primary key (id)
) CHARACTER SET utf8 COLLATE utf8_general_ci comment '其他奖品数据表';


```

## 中奖纪录
```sql

create table gift_record
(
	id int auto_increment comment '其他奖品的ID，主键自增',
	gift_id int not null comment '奖品ID - 关联gift表中的ID',
	gift_name varchar(255) not null comment '奖品名称',
	gift_type int not null comment '奖品Type - 关联gift表中的gtype',
	uid int not null comment '用户ID - 发放给谁了',
	username varchar(255) not null comment '用户Nick',
    prize_code int not null comment '用户抽奖时的code码',
    gift_data varchar(255) comment '获奖说明',
    sys_status tinyint not null comment '状态 0正常 1删除 2异常',
	sys_created int null comment '创建时间',
	sys_ip varchar(50) null comment '用户的IP',
	constraint prize_pk
		primary key (id)
) CHARACTER SET utf8 COLLATE utf8_general_ci comment '获奖纪录';

```
## 黑名单-用户
```sql
create table user_black
(
	id int auto_increment comment '用户ID',
	username varchar(255) not null comment '用户Nick',
	blacktime int not null comment '黑名单时间',
	realname varchar(50)  comment '联系人',
	mobile varchar(50) comment '用户手机号',
	address varchar(255) comment '地址',
	sys_update int null comment '更新时间',
	sys_created int null comment '创建时间',
	sys_ip varchar(50) null comment '用户的IP',
	constraint prize_pk
		primary key (id)
) CHARACTER SET utf8 COLLATE utf8_general_ci comment '用户黑名单';

```

## IP黑名单 - IP
```sql
create table ip_black
(
	id int auto_increment comment 'IP黑名单的ID',
	ip varchar(255) not null comment 'ID地址',
	blacktime int not null comment '黑名单时间',
	sys_update int null comment '更新时间',
	sys_created int null comment '创建时间',
	constraint prize_pk
		primary key (id)
) CHARACTER SET utf8 COLLATE utf8_general_ci comment 'IP黑名单';

```

## 用户参与次数
```sql
create table count_record
(
	id int auto_increment comment '每日抽奖次数表的ID',
	uid int not null comment '用户ID',
	day int not null comment '日期',
	num int null comment '抽奖次数',
	sys_created int null comment '创建时间',
	sys_update int null comment '更新时间',
	constraint prize_pk
		primary key (id)
) CHARACTER SET utf8 COLLATE utf8_general_ci comment '每日次数表';

```