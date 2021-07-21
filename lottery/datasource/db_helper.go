package datasource

import (
	"com.wangzhumo.lottery/conf"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
	"xorm.io/xorm"
)

// 为了避免并发造成问题，加入互斥锁
var dbLock sync.Mutex


// 实例中的Engine
var masterInstance *xorm.Engine

// InstanceDB 获取实例
func InstanceDB() *xorm.Engine{
	// 有就直接返回
	if masterInstance != nil {
		return masterInstance
	}

	// 锁只是为了避免重复创建数据库Engine
	dbLock.Lock()
	defer dbLock.Unlock()
	// 避免重复
	if masterInstance != nil {
		return masterInstance
	}

	return NewInstance()
}


//NewInstance root:wzmmysql@tcp(127.0.0.1:3306)/lottery?charset=utf8
func NewInstance() *xorm.Engine {
	mysqlConnConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		conf.DbMaster.User,
		conf.DbMaster.Psd,
		conf.DbMaster.Host,
		conf.DbMaster.Port,
		conf.DbMaster.Database)
	engine, err := xorm.NewEngine(conf.DriverName, mysqlConnConfig)
	if err != nil {
		log.Fatal("mysql can't instance,err",err)
		return nil
	}
	masterInstance = engine
	masterInstance.ShowSQL(true)
	return masterInstance
}
