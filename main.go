package main

import (
	//"github.com/gin-gonic/gin"
	"zixu.com/cmdAdmin/conf"
	"zixu.com/cmdAdmin/conf/vars"
	"zixu.com/cmdAdmin/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()

	r.Run(vars.ServerAddr)
}
