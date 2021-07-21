package conf


type RedisConfig struct {
	Host string
	Port string
	User string
	Psd string
	IsRunning bool
}


var RedisConfigList = []RedisConfig{
	{
		Host: "127.0.0.1",
		Port: "6379",
		User: "",
		Psd: "wzmredis",
		IsRunning: true,
	},
}

var RedisMaster RedisConfig = RedisConfigList[0]

//defaultDb = root:wzmmysql@tcp(127.0.0.1:3306)/lottery?charset=utf8
//redisdb=127.0.0.1:6379
//redispsd=wzmredis
//rabbitmqdb=amqp://wzmrabbit:wzmrabbit@127.0.0.1:5672/