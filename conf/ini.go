package conf

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
	"time"
	"zixu.com/cmdAdmin/conf/vars"
)

const timeOut = 10

func SetConfig(filename string, handle func(cfgFile *ini.File), ) {
	go Config(filename, handle,time.Second * timeOut )
}
func Config(filename string, handle func(cfgFile *ini.File),timeout time.Duration ) {
	for {
		cfgFile, err := ini.Load(filename)
		if err != nil{
			fmt.Println(err)
		}else {
			handle(cfgFile)
		}
		time.Sleep(timeout)

	}
}

func handle(cfgFile *ini.File){
	db := cfgFile.Section("db")
	DB := vars.DB
	DB.Addr = db.Key("addr").MustString(DB.Addr)
	DB.User = db.Key("user").MustString(DB.User)
	DB.Password = db.Key("password").MustString(DB.Password)
	DB.DataBase = db.Key("database").MustString(DB.DataBase)
	redis := cfgFile.Section("redis")
	Redis := vars.Redis
	Redis.Used = redis.Key("used").MustBool(Redis.Used)
	Redis.Addr = redis.Key("addr").MustString(Redis.Addr)
	Redis.Password = redis.Key("password").MustString(Redis.Password)
	Redis.DB = redis.Key("db").MustInt(Redis.DB)
	vars.LogLevel = cfgFile.Section("log").Key("level").MustString(vars.LogLevel)
	vars.GinMode = cfgFile.Section("log").Key("gin_mode").MustString(vars.GinMode)
	vars.SessionSecret = cfgFile.Section("server").Key("session_secret").MustString(vars.SessionSecret)
	vars.ServerAddr = cfgFile.Section("server").Key("addr").MustString(vars.ServerAddr)

}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func init()  {
	filename := "conf.ini"
	args :=os.Args[1:]
	if len(args) >0{
		filename = args[0]

	}

	cfgFile, err := ini.Load(filename)
	panicErr(err)
	handle(cfgFile)
}
