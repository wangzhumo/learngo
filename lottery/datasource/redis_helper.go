package datasource

import (
	"com.wangzhumo.lottery/conf"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"sync"
	"time"
)

// 为了避免并发造成问题，加入互斥锁
var redisLock sync.Mutex


// 实例中的Engine
var redisInstance *RedisConn

// InstanceRedis 获取实例
func InstanceRedis() *RedisConn{
	// 有就直接返回
	if redisInstance != nil {
		return redisInstance
	}

	// 锁只是为了避免重复创建数据库Engine
	redisLock.Lock()
	defer redisLock.Unlock()
	// 避免重复
	if redisInstance != nil {
		return redisInstance
	}

	return NewRedisInstance()
}


type RedisConn struct {
	pools *redis.Pool
	isDebug bool
}


// Do 执行Redis操作
func (redisConn *RedisConn) Do(command string,args ...interface{}) (
	reply interface{},err error) {
	// 获取一个连接
	conn := redisConn.pools.Get()
	defer conn.Close()
	// 纪录
	timestampStart := time.Now().Unix()
	reply, err = conn.Do(command, args...)
	if err != nil {
		// 获取信息并打印
		errMsg := conn.Err()
		if errMsg != nil {
			log.Println("redis do err:",errMsg)
		}
	}
	timestampEnd  := time.Now().Unix()
	// 如果需要则打印
	if redisConn.isDebug {
		fmt.Printf("[Redis Info] [%dus]cmd=%s , err=%s , args=%v , reply=%s\n",
			(timestampEnd-timestampStart)/1000,command,err,args,reply)
	}
	return
}

// Debug 设置是否Debug
func (redisConn *RedisConn) Debug(debug bool)  {
	redisConn.isDebug = debug
}

// NewRedisInstance 创建Redis
func NewRedisInstance() *RedisConn {
	connect := PoolConnect()
	instance := &RedisConn{
		pools: connect,
	}
	redisInstance = instance
	redisInstance.isDebug = true
	return redisInstance
}

// PoolConnect 获取Redis的链接
func PoolConnect() *redis.Pool {
	redisPsd := redis.DialPassword(conf.RedisMaster.Psd)
	host := fmt.Sprintf("%s:%s",
		conf.RedisMaster.Host,conf.RedisMaster.Port)
	redisPool := &redis.Pool{
		MaxIdle:     100,
		MaxActive:   100,
		IdleTimeout: 100 * time.Second,
		Wait:        true,
		MaxConnLifetime: 0,
		Dial: func() (redis.Conn, error) {
			dial, err := redis.Dial("tcp", host, redisPsd)
			if err != nil {
				log.Fatal("redis connect fail err:",err)
				return nil, err
			}
			//if _,err := dial.Do("AUTH",redisPsd); err != nil {
			//	err := dial.Close()
			//	if err != nil {
			//		return nil, err
			//	}
			//	return nil, err
			//}
			return dial, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	return redisPool
}


