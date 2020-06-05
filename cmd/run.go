package cmd

import (
	"bufio"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"os"
	"os/exec"
	"syscall"
	"zixu.com/cmdAdmin/model"
)

const (
	ProgramNotRun = 0
	ProgramStart  = 1
	ProgramEnd    = 2
)

type Program struct {
	Name      string
	Arg       []string
	OutChan   chan string
	ErrChan   chan string
	EndChan   chan bool
	errStdout io.ReadCloser
	stdout    io.ReadCloser
	Process   *os.Process
	Cmd       *exec.Cmd
	Dir       string
	State     int
}

func (p *Program) Run() error {
	cmd := exec.Command(p.Name, p.Arg...)
	cmd.Dir = p.Dir
	//exec.CommandContext()
	//fmt.Println(cmd.Path, cmd.Dir)
	stdout, err := cmd.StdoutPipe()
	if err != nil { //获取输出对象，可以从该对象中读取输出结果
		return err
	}
	errStdout, err := cmd.StderrPipe()
	if err != nil { //获取输出对象，可以从该对象中读取输出结果
		return err
	}
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	if err := cmd.Start(); err != nil { // 运行命令
		return err
	}
	var outChan = make(chan string)
	var ErrChan = make(chan string)
	p.EndChan = make(chan bool, 10)
	go p.OutStdout(outChan, stdout)
	go p.OutStdout(ErrChan, errStdout)
	p.OutChan = outChan
	p.ErrChan = ErrChan
	p.errStdout = errStdout
	p.stdout = stdout
	p.Cmd = cmd
	p.State = 1
	p.Process = cmd.Process
	AddMap(p)
	return nil
}

func (p *Program) WriteFile(path string) error {
	file, err := os.OpenFile(path, os.O_CREATE, os.ModePerm)
	if err != nil {
		p.Close()
		return err
	}
	go func(p *Program, file *os.File) {
		defer file.Close()
		for {
			select {
			case d := <-p.ErrChan:
				_, _ = file.WriteString(d + "\n")
			case d := <-p.OutChan:
				_, _ = file.WriteString(d + "\n")
			case <-p.EndChan:
				return
			}
		}
	}(p, file)
	return nil
}
func (p *Program)WriteData(task *model.Task) {
	go func(p *Program,task *model.Task) {
		endIndex := 2
		for {
			select {
			case d := <-p.ErrChan:
				if d != "\n"{
					data :=model.TaskOutput{TaskID: task.ID,Type:model.ErrOutput,Text:d}
					model.GetDB().Create(&data)
				}
			case d := <-p.OutChan:
				if d != "\n"{
					data :=model.TaskOutput{TaskID: task.ID,Type:model.Output,Text:d}
					model.GetDB().Create(&data)
				}

			case <-p.EndChan:
				endIndex --
				if endIndex ==0{
					fmt.Println("end")
					return
				}
			}
		}
	}(p,task)
}

func (p *Program) Close() {
	if p.State == 2 {
		return
	}
	p.State = 2

	close(p.OutChan)
	close(p.ErrChan)
	p.stdout.Close()
	p.errStdout.Close()

	p.EndChan <- true
	p.EndChan <- true

}


func (p *Program) SafeClose() {
	p.OutChan <- "\n"
	p.ErrChan <- "\n"
	p.State = 2
	p.EndChan <- true
}

func ConvertByte2String(str string) string {
	var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().String(str)
	return decodeBytes
}
func (p *Program) OutStdout(ch chan string, stdout io.ReadCloser) {
	//defer close(ch)
	defer func() {
		_ = recover()
	}()
	scan := bufio.NewScanner(stdout)
	for scan.Scan() {
		data :=ConvertByte2String(scan.Text())
		ch <- data
	}
	p.SafeClose()
	defer stdout.Close()

}

func (p *Program) Kill()error{

	c :=exec.Command("taskkill","-F","-T","/PID",fmt.Sprintf("%d",p.Process.Pid))
	return c.Run()
}
