package server

import (
	"github.com/gin-gonic/gin"
	"zixu.com/cmdAdmin/conf/vars"
	"zixu.com/cmdAdmin/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	// 中间件, 顺序不能改
	gin.SetMode(vars.GinMode)

	r.Use(middleware.Session(vars.SessionSecret))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())
	// 路由
	system(r)
	task(r)
	r.Static("/html","static")
	r.GET("/", func(c *gin.Context) {
		c.Redirect(301,"/html")
	})
	return r
}
