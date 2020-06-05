package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Name    string
	Dir     string `gorm:"size:1024"`
	Command string `gorm:"size:4096"`
	//1定时重启,2挂掉重启
	KeepType int64
	//定时重启周期
	Timing int64
	//下次重启时间
	NextKeepTime time.Time `gorm:"default:null"`
	Remark  string `gorm:"size:4096"`
	PID int `gorm:"column:pid"`
}

func (t *Task) GetOutput() []TaskOutput {
	var tops []TaskOutput
	DB.Where("task_id=?", t.ID).Find(&tops)
	return tops
}


func (t *Task) DeleteOutput()  {
	top:= TaskOutput{}
	DB.Where("task_id=?", t.ID).Delete(&top)

}

const (
	ErrOutput = 1
	Output = 2
)
type TaskOutput struct {
	gorm.Model
	TaskID uint
	Type   int8
	Text   string `gorm:"size:4096"`
}
