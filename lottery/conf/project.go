package conf

import "time"

const SysTimeFormat = "2006-01-02 15:04:05"
const SysTimeFormatShort = "2006-01-02"

var SysTimeLocation, _ = time.LoadLocation("Asia/Shanghai")

var SignSecret = []byte("oe933d1sh7")

var CookieSecret = "wzmlottery"
