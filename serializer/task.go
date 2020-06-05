package serializer

import (
	"zixu.com/cmdAdmin/cmd"
	"zixu.com/cmdAdmin/model"
)

type Task struct {
	model.Task
	State int

}
func (t *Task)Bind(task * model.Task){
	t.Task = *task
	program := cmd.GetProgram(task.PID)
	//fmt.Println(program)
	if program != nil{
		t.State = program.State
	}
}


