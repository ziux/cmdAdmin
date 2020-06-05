package server

import (
	"github.com/gin-gonic/gin"
	"zixu.com/cmdAdmin/api"
	"zixu.com/cmdAdmin/middleware"
)

// NewRouter 路由配置
func system(r *gin.Engine){
	v1 := r.Group("/api/system")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
		}
	}
}
