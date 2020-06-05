package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"zixu.com/cmdAdmin/model"
	"zixu.com/cmdAdmin/serializer"
	"zixu.com/cmdAdmin/service"
)

func TaskCreate(c *gin.Context) {
	var s service.TaskService
	if err := c.ShouldBind(&s); err == nil {
		res := s.Create(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func TaskUpdate(c *gin.Context) {
	task := getTask(c)
	var s service.TaskService
	if err := c.ShouldBind(&s); err == nil {
		res := s.Update(task,c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func TaskList(c *gin.Context) {
	db := model.GetDB()
	var queryset []model.Task
	dataList := make([]serializer.Task,0)
	db.Find(&queryset)
	for i:=0;i<len(queryset);i++{
		task := serializer.Task{}
		task.Bind(&queryset[i])
		dataList = append(dataList, task)
	}
	c.JSON(200, serializer.Response{Data: dataList})
}

func TaskDelete(c *gin.Context) {
	task := getTask(c)
	db := model.GetDB()
	db.Delete(task)
	c.JSON(200, serializer.Response{Data: task})
}

func TaskKill(c *gin.Context) {
	task := getTask(c)
	err := service.Kill(task)
	if err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}
	c.JSON(200, serializer.Response{Data: task})
}

func TaskStart(c *gin.Context) {
	task := getTask(c)
	program, err := service.Run(task)
	if err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}
	c.JSON(200, serializer.Response{Data: program.Process.Pid})
}


func TaskOutput(c *gin.Context){
	task := getTask(c)
	dataList := task.GetOutput()
	c.JSON(200, serializer.Response{Data: dataList})
}

func DeleteOutput(c *gin.Context){
	task := getTask(c)
	task.DeleteOutput()
	c.JSON(200, serializer.Response{Data: "ok"})
}

func getTask(c *gin.Context) *model.Task {
	id := c.Query("id")
	Iid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(200, ErrorResponse(err))
		return nil
	}
	db := model.GetDB()
	var task model.Task
	db.Where("id=?", Iid).First(&task)
	if task.ID == 0 {
		c.JSON(200, ErrorResponse(errors.New("id错误")))
		return nil
	}
	return &task
}
