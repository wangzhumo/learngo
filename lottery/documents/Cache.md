# 缓存设计

## 设计依据

- 目标：提高系统的性能，减少数据库依赖
- 原则：平衡“性能，开发周期，复杂度”
- 方向：数据读多写少，数据量有限，数据分散（结构关系性弱）

## Redis

- 奖品         数量少，修改少，但是访问量大 ->  可以全量缓存
- 优惠券    数量有限，修改少 ->  可以全量缓存
- 中奖纪录        读写量一般，但是数据量较大 -> 部分缓存（统计，一段时间内的纪录）  
- 黑名单IP/UID     读多写少 
- 参与次数     收益不大



