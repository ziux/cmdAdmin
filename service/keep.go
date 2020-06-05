package service

import (
	"fmt"
	"time"
	"zixu.com/cmdAdmin/cmd"
	"zixu.com/cmdAdmin/model"
)

func Range()  {
	for {
		time.Sleep(10 * time.Second)
		db := model.GetDB()
		var taskList []model.Task
		db.Where("keep_type in (1,2)").Find(&taskList)
		for i:=0;i<len(taskList);i++{
			task := taskList[i]
			switch task.KeepType {
			case 1:
				Reload(&task)
			case 2:
				KeepRun(&task)
			}




		}
	}
}
func init()  {
	go Range()
}
//定时重启
func Reload(task *model.Task){
	if task.NextKeepTime.Unix() < time.Now().Unix(){
		_ = Kill(task)
		_, _ = Run(task)
		task.NextKeepTime =  time.Now().Add(time.Second * time.Duration(task.Timing))
		fmt.Println("Reload",task)
		model.GetDB().Save(&task)
	}
}
func KeepRun(task *model.Task){
	if task.PID !=0 {
		if program :=cmd.GetProgram(task.PID); program!= nil  {
			if program.State==1 {
				return
			}else {
				cmd.DelMap(task.PID)
			}
		}
	}
	fmt.Println("KeepRun",task)
	_, _ = Run(task)

}
