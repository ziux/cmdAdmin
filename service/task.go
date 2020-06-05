package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"os/exec"
	"time"
	"zixu.com/cmdAdmin/cmd"
	"zixu.com/cmdAdmin/conf"
	"zixu.com/cmdAdmin/model"
	"zixu.com/cmdAdmin/serializer"
	"zixu.com/cmdAdmin/util"
)

type TaskService struct {
	Name    string `json:"name"`
	Dir     string `json:"dir"`
	Command string `json:"command"`
	Remark  string `json:"remark"`
	KeepType int64 `json:"keepType"`
	Timing int64`json:"timing"`

}

func (t *TaskService) Create(c *gin.Context) serializer.Response {
	err := t.Verification()
	if err != nil {
		return serializer.ParamErr("", err)
	}
	var task = model.Task{Name: t.Name, Dir: t.Dir, Command: t.Command, Remark: t.Remark,
		KeepType:t.KeepType,Timing: t.Timing,NextKeepTime: time.Now()}
	db := model.GetDB()
	db.Create(&task)
	//_,err = Run(&task)
	//if err != nil {
	//	return serializer.ParamErr("", err)
	//}
	return serializer.Response{Data: task}
}

func (t *TaskService) Update(task *model.Task,c *gin.Context) serializer.Response {
	task.Dir = t.Dir
	task.Command = t.Command
	task.Remark = t.Remark
	task.KeepType = t.KeepType
	task.Timing = t.Timing
	task.NextKeepTime = time.Now()
	db := model.GetDB()
	db.Save(&task)
	return serializer.Response{Data: task}
}





func Run(t *model.Task) (*cmd.Program,error){
	if t.PID !=0 {
		if program :=cmd.GetProgram(t.PID); program!= nil  {
			if program.State==1 {
				return nil, errors.New("task_has_run")
			}else {
				cmd.DelMap(t.PID)
			}

		}
	}
	name, args := util.ParseCommand(t.Command)
	program := &cmd.Program{
		Name: name,
		Arg:  args,
		Dir:  t.Dir,
	}
	err := program.Run()
	if err != nil {
		return nil, err
	}
	program.WriteData(t)
	t.PID = program.Process.Pid
	model.GetDB().Save(t)
	return program, nil

}

func Kill(t *model.Task) error{
	program  := cmd.GetProgram(t.PID)
	if program == nil{
		return errors.New("no_program")
	}
	//return program.Process.Signal(syscall.SIGQUIT)
	//return syscall.Kill(-program.Process.Pid, syscall.SIGKILL)
	c :=exec.Command("taskkill","-F","-T","/PID",fmt.Sprintf("%d",program.Process.Pid))
	return c.Run()
	//return program.Process.Kill()
}

func (t *TaskService) Verification() error {
	if t.Name == "" || t.Command == "" {
		return errors.New(conf.T("task_command_name_verification"))
	}
	db := model.GetDB()
	var task model.Task
	db.Where("name=?", t.Name).First(&task)
	if task.ID != 0 {
		return errors.New(conf.T("task_already_exists"))
	}
	return nil
}
