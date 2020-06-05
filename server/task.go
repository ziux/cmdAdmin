package server

import (
	"github.com/gin-gonic/gin"
	"zixu.com/cmdAdmin/api"
)

func task(r *gin.Engine){
	task := r.Group("/api/task")
	//task.Use(middleware.AuthRequired())
	task.POST("",api.TaskCreate)
	task.PUT("",api.TaskUpdate)
	task.GET("",api.TaskList)
	task.DELETE("",api.TaskDelete)
	task.GET("/start",api.TaskStart)
	task.GET("/kill",api.TaskKill)
	task.GET("/output",api.TaskOutput)
	task.GET("/output/delete",api.DeleteOutput)


}
