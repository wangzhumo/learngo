package conf

const DriverName = "mysql"

type DbConfig struct {
	Host string
	Port string
	User string
	Psd string
	Database string
	IsRunning bool
}


var DbConfigList = []DbConfig{
	{
		Host: "127.0.0.1",
		Port: "3306",
		User: "root",
		Psd: "wzmmysql",
		Database: "lottery",
		IsRunning: true,
	},
}

var DbMaster DbConfig = DbConfigList[0]

//defaultDb = root:wzmmysql@tcp(127.0.0.1:3306)/lottery?charset=utf8
//redisdb=127.0.0.1:6379
//redispsd=wzmredis
//rabbitmqdb=amqp://wzmrabbit:wzmrabbit@127.0.0.1:5672/